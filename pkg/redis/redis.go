package redis

import (
	"fmt"
	"server-aggregation/config"
	"sync"

	"github.com/go-redis/redis"
)

var once sync.Once
var redisClient *redis.Client

const Nil = redis.Nil

var redisMap map[string]*redis.Client

func Init() {
	once.Do(func() {
		redisMap = make(map[string]*redis.Client)
		redisConfigs := config.GetStringMap("redis")
		for redisName, _ := range redisConfigs {
			redisClient = redis.NewClient(&redis.Options{
				Addr:       config.GetString(fmt.Sprintf("redis.%s.addr", redisName)),
				Password:   config.GetString(fmt.Sprintf("redis.%s.password", redisName)), // no password set
				DB:         config.GetInt(fmt.Sprintf("redis.%s.db", redisName)),          // use default DB
				PoolSize:   config.GetInt(fmt.Sprintf("redis.%s.pool", redisName)),
				MaxRetries: 3,
			})

			pong, err := redisClient.Ping().Result()
			if err == nil {
				fmt.Printf("\033[1;30;42m[info]\033[0m redis: %s, connect success %s\n", redisName, pong)
			} else {
				panic(fmt.Sprintf("\033[1;30;41m[error]\033[0m redis: %s,  connect error %s\n", redisName, err.Error()))
			}
			redisMap[redisName] = redisClient
		}
	})
}

// GetBaseRedis 固件检测的是 5788
func GetBaseRedis() *redis.Client {
	return redisMap["base"]
}

// GetBase2Redis 插件化任务和 固件入库任务是 8258
func GetBase2Redis() *redis.Client {
	return redisMap["base2"]
}
