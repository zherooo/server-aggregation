package router

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"server-aggregation/config"
	"server-aggregation/pkg/validate"
	"time"

	"github.com/gin-gonic/gin/binding"
)

func Run(handler *gin.Engine, addr string) {
	binding.Validator = validate.GinValidator()

	gin.SetMode(config.CfgEnv)

	srv := &http.Server{
		Addr:    addr,
		Handler: handler,
		//ReadTimeout:    20,
		//WriteTimeout:   20,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Printf("\033[1;30;42m[info]\033[0m start http server listening %s\n", addr)
	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("\033[1;30;41m[error]\033[0m start http server error: ", err.Error())
			os.Exit(1)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Println("\n\033[1;30;42m[info]\033[0m Shutdown Server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Printf("\033[1;30;42m[info]\033[0m Server Shutdown: %s", err)
	}
	fmt.Println("Server exited")
}
