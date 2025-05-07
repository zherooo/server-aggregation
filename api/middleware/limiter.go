package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	"go.uber.org/zap"

	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	sredis "github.com/ulule/limiter/v3/drivers/store/redis"
)

const LimiterLogNamed = "limiter"

// 处理跨域请求,支持options访问
// rate: 1-M,1-H,1-D| 次数,
func Limiter(logger *zap.Logger, rate string) gin.HandlerFunc {
	// Define a limit rate to 4 requests per hour.
	rateFormatted, err := limiter.NewRateFromFormatted(rate)
	if err != nil {
		logger.Error("limiter rate error", zap.Error(err))
	}

	// Create a store with the redis client.
	store, err := sredis.NewStoreWithOptions(nil, limiter.StoreOptions{
		Prefix:   "ft:limiter:",
		MaxRetry: 3,
	})
	if err != nil {
		logger.Error("limiter store redis error", zap.Error(err))
	}
	// Create a new middleware with the limiter instance.
	return mgin.NewMiddleware(limiter.New(store, rateFormatted))
}
