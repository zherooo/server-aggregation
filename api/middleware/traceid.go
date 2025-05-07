package middleware

import (
	"github.com/gin-gonic/gin"
	"server-aggregation/pkg/log"
	"server-aggregation/pkg/sonyflake"
	"strconv"
)

// context中添加trace_id字段
func AddTraceId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 处理请求
		uniqueID, _ := sonyflake.ID()
		ctx.Set(log.TraceID, strconv.FormatUint(uniqueID, 10))
		ctx.Next()
	}
}
