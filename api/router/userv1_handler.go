package router

import (
	"fmt"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"server-aggregation/api/controller/health"
	"server-aggregation/api/middleware"
	"server-aggregation/config"
	"server-aggregation/pkg/log"
)

var handle *gin.Engine

func StartUserV1Route() {
	addr := fmt.Sprintf(":%d", config.GetInt("app.user_port"))
	Run(StartUserV1Handler(), addr)
}

func StartUserV1Handler() *gin.Engine {
	handle = gin.New()
	handle.ForwardedByClientIP = true

	// 探活接口
	handle.GET("/", health.Hello)
	handle.HEAD("/health", health.Hello)
	handle.GET("/health", health.Hello)
	handle.GET("/ping", health.Ping)
	// 正式环境接收 Recover 日志
	//handle.Use(middleware.RecoveryWithZap(log.New().Named(middleware.RecoveryLogNamed), true))

	// 开启 gzip
	handle.Use(gzip.Gzip(gzip.DefaultCompression))

	// context中添加 trace_id
	handle.Use(middleware.AddTraceId())

	// 根据配置决定是否启用 api 请求日志
	if config.GetBool("log.request_log") {
		handle.Use(middleware.WriterLog(log.New().Named(middleware.RequestLogNamed)))
	}

	return handle
}
