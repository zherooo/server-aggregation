package consts

const (
	Config    = iota + 1 // 配置文件
	Logger               // 日志
	Mysql                // mysql 资源
	Redis                // redis  资源
	MongoDB              // mongodb 资源
	Docker               // docker daemon  资源
	Task                 // task 服务的cron
	UserV1API            // user api 服务
)
