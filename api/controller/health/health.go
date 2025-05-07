package health

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Hello 首页，用于健康检查
func Hello(c *gin.Context) {
	c.String(http.StatusOK, "hello, word.")
}

func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
