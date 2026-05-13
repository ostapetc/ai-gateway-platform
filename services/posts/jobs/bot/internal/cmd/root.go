package cmd

import (
	"os"

	"github.com/ostapetc/ai-gateway-platform/services/posts/jobs/bot/internal/config"
	"github.com/ostapetc/ai-gateway-platform/services/posts/jobs/bot/internal/logic"
	"github.com/ostapetc/ai-gateway-platform/services/posts/jobs/bot/internal/svc"
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

	printTimeJob = &cobra.Command{
		Use:   "printtime",
		Short: "print current time",
		RunE:  logic.PrintTime,
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&confPath, "config", "etc/cron.yaml", "config file")
	rootCmd.AddCommand(printTimeJob)
}

func initConfig() {
	var c config.Config
	conf.MustLoad(confPath, &c)
	svc.InitSvcCtx(c)
}
