package cmd

import (
	"os"

	"github.com/ostapetc/ai-gateway-platform/services/users-bot-cronjob/internal/config"
	"github.com/ostapetc/ai-gateway-platform/services/users-bot-cronjob/internal/logic"
	"github.com/ostapetc/ai-gateway-platform/services/users-bot-cronjob/internal/svc"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/conf"
)

const codeFailure = 1

var (
	confPath string

	rootCmd = &cobra.Command{
		Use:   "bot",
		Short: "exec cron job",
	}

	createUsersJob = &cobra.Command{
		Use:   "createusers",
		Short: "create a random user via users gRPC service",
		RunE:  logic.CreateUsers,
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(codeFailure)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&confPath, "config", "etc/cron.yaml", "config file")
	rootCmd.AddCommand(createUsersJob)
}

func initConfig() {
	var c config.Config
	conf.MustLoad(confPath, &c, conf.UseEnv())
	svc.InitSvcCtx(c)
}
