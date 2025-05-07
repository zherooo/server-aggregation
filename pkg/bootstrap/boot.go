package bootstrap

import (
	"server-aggregation/api/router"
	"server-aggregation/config"
	"server-aggregation/internal/consts"
	"server-aggregation/internal/cron"
	"server-aggregation/internal/model/mongodb"
	"server-aggregation/internal/model/mysql"
	"server-aggregation/pkg/docker"
	"server-aggregation/pkg/log"
	"server-aggregation/pkg/redis"
)

var initFuncMap = map[int]func(){
	consts.Config:    config.Init,
	consts.Logger:    log.Init,
	consts.Mysql:     mysql.Init,
	consts.Redis:     redis.Init,
	consts.MongoDB:   mongodb.Init,
	consts.Docker:    docker.Init,
	consts.Task:      cron.InitTask,
	consts.UserV1API: router.StartUserV1Route,
}

func Init(resourceType ...int) {
	for _, item := range resourceType {
		if bootFunc, ok := initFuncMap[item]; ok {
			bootFunc()
		}
	}
}
