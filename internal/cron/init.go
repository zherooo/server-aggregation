package cron

import (
	"os"
	"server-aggregation/pkg/async"
)

var Client *async.Async

func InitTask() {
	ch := make(chan os.Signal)

	Client = async.NewAsync(ch)
	// 接受信息
	Client.Register(HandleUrlResult)
}
