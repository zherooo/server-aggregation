package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"server-aggregation/cmd/task"
	"server-aggregation/cmd/userv1"
	"server-aggregation/config"
	"server-aggregation/pkg/utils"

	"github.com/spf13/cobra"
)

// RootCmd RootCmd
var RootCmd = &cobra.Command{
	Use:              "server-aggregation",
	Short:            "基本集成框架",
	Long:             "基本集成框架聚合服务项目",
	TraverseChildren: true,
	CompletionOptions: cobra.CompletionOptions{
		HiddenDefaultCmd: true,
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		// 初始化依赖扩展
		validate()
	},
}

// validate 参数校验
func validate() {
	if !utils.InArray(config.CfgEnv, []string{gin.DebugMode, gin.TestMode, gin.ReleaseMode}) {
		fmt.Println("[ERROR] --env 环境错误，请检查启动参数...")
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringVar(&config.CfgEnv, "env", "test", "环境配置 (debug本地开发环境、test测试环境、release生产环境)")
	RootCmd.AddCommand(userv1.ServerCmdUserV1)
	RootCmd.AddCommand(task.ServerCmdTask)
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
