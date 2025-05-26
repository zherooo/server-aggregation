## 关于

> 本项目基于gin封装了常用的组件
> 
> 根据不同启动命令使用不同的服务


### 开发环境
golang 2.24.2+

### 目录结构
```text
.
├── Dockerfile
├── README.md
├── api
│   ├── controller // 控制器controller
│   ├── middleware // 中间件
│   ├── router	 // 路由
│   └── server.go
├── cmd    // 命令行cmd
├── config // 配置文件
├── go.mod
├── go.sum
├── internal
│   ├── consts  // 项目涉及的通用常量
│   ├── model   // 数据库资源
│   └── service // 服务 service
├── logs // 日志目录
├── pkg  // 内部组件（日志、配置等）
├── main.go
├── makefile
└── tests // 测试文件目录
```

### 配置文件：
> 通过启动参数 `--env` 更改运行环境

- --env debug，表示设置本地开发环境，使用 configs/debug_configs.yaml
- --env test，表示设置为测试环境，使用 configs/test_configs.yaml
- --env release，表示设置为生产环境/私有化部署，使用 configs/release_configs.yaml

### 使用方法
安装扩展： 
``` shell
go mod tidy
```

编译程序：
``` makefile
make build
```

项目启动示例
```shell
./server-aggregation [module_name] --env [environment]
```
本地调试前置条件
````
创建配置文件中名为 userv,task 数据库
修改配置文件中redis密码
````
本地调试测试：
```shell
go run main.go userv1 --env debug
```

docker-compose.yml中，可根据command不同，使用不同的服务
```
#web服务
  userv1-web:
    image: ${aggregation_image}
    hostname: "userv1-web"
    container_name: "userv1-web"
    environment:
      HOST_ADDR: http://${IP_ADDR}
      IP_ADDR: ${IP_ADDR}
    command: /opt/server-aggregation userv1 --env debug
    privileged: true
    restart: always
#任务服务
  task:
    image: ${aggregation_image}
    hostname: "task"
    container_name: "task"
    command: /opt/server-aggregation task --env test
```

### 集成组件：
* 支持 [cobra](https://github.com/spf13/cobra) cli工具
* 支持 [zap](https://go.uber.org/zap) 日志收集
* 支持 [viper](https://github.com/spf13/viper) 配置文件解析
* 支持 [gorm](https://gorm.io/gorm) 数据库组件
* 支持 [go-redis](https://github.com/go-redis/redis/v7) redis组件
* 支持 [cron](https://github.com/robfig/cron) 定时任务
* 支持 [MongoDB](https://go.mongodb.org/mongo-driver) MongoDB 组件
* 支持 RESTful API 返回值规范

