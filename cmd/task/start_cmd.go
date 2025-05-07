package task

import (
	"github.com/spf13/cobra"
	"server-aggregation/internal/consts"
	"server-aggregation/pkg/bootstrap"
)

// ServerCmdTask server
var ServerCmdTask = &cobra.Command{
	Use:   "task",
	Short: "task任务项目",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		// 资源初始化
		bootstrap.Init(
			consts.Config,
			consts.Logger,
			consts.Mysql,
			consts.Task,
			consts.UserV1API,
		)
	},
}
