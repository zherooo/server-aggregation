package middleware

import (
	"bytes"
	"io"
	"net/http"
	"server-aggregation/pkg/log"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const RequestLogNamed = "http_request"
const maxLogSize = 2 * 1024 * 1024 // 2 MB 日志记录最大的值

// WriterLog 处理跨域请求,支持options访问
func WriterLog(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// 准备相应日志
		bodyBuf := new(bytes.Buffer)
		// 通过httptest方式调用Get请求时c.Request.Body为nil
		if c.Request.Body != nil {
			_, _ = io.Copy(bodyBuf, c.Request.Body)
		}
		body := bodyBuf.Bytes()
		c.Request.Body = io.NopCloser(bytes.NewReader(body))
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		c.Next()

		// 请求的body
		var requestBody string
		if len(body) > maxLogSize {
			requestBody = "request body too large, not logged"
		} else {
			requestBody = string(body)
		}

		// 相应的body
		var respBody string
		if blw.body.Len() > maxLogSize {
			respBody = "response body too large, not logged"
		} else {
			respBody = blw.body.String()
		}

		latency := time.Since(start)
		var fs = []zap.Field{
			zap.Int("status", c.Writer.Status()),
			zap.String("ip", c.ClientIP()),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.String("query", c.Request.URL.RawQuery),
			zap.Duration("latency", latency),
			zap.String("resp", respBody),
			zap.String("agent", c.Request.UserAgent()),
			zap.Any("session", nil), // 信息太大暂时用不到
		}
		// Append error field if this is an erroneous req.
		if len(c.Errors) > 0 {
			fs = append(fs, zap.Strings("errors", c.Errors.Errors()))
		}
		// Post Method writer body
		if c.Request.Method != http.MethodGet {
			fs = append(fs, zap.String("body", requestBody))
		}
		// Writer X-Request-ID to log
		xRequestId := c.Request.Header.Get("X-Request-ID")
		if len(xRequestId) > 0 {
			fs = append(fs, zap.String("request_id", xRequestId))
		}
		// Writer trace_id to log
		if traceID := c.Value(log.TraceID); traceID != nil {
			fs = append(fs, zap.String(log.TraceID, traceID.(string)))
		}
		if driverTrace := c.GetHeader("X-Driver-Request-ID"); driverTrace != "" {
			fs = append(fs, zap.String("driver_client_trace", driverTrace))
		}
		logger.Info(c.Request.RequestURI, fs...)
	}
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
