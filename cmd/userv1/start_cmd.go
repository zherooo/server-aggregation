package userv1

import (
	"github.com/spf13/cobra"
	"server-aggregation/internal/consts"
	"server-aggregation/pkg/bootstrap"
)

// ServerCmdUserV1 http server
var ServerCmdUserV1 = &cobra.Command{
	Use:   "userv1",
	Short: "用户模块",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		// 资源初始化
		bootstrap.Init(
			consts.Config,
			consts.Logger,
			consts.Mysql,
			consts.Redis,
			consts.MongoDB,
			consts.UserV1API,
		)
	},
}
