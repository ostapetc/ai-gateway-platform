package cmd

import (
	"os"

	"github.com/ostapetc/ai-gateway-platform/services/posts-bot-cronjob/internal/config"
	"github.com/ostapetc/ai-gateway-platform/services/posts-bot-cronjob/internal/logic"
	"github.com/ostapetc/ai-gateway-platform/services/posts-bot-cronjob/internal/svc"
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

	createPostsJob = &cobra.Command{
		Use:   "createposts",
		Short: "create a random post via posts gRPC service",
		RunE:  logic.CreatePosts,
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
	rootCmd.AddCommand(createPostsJob)
}

func initConfig() {
	var c config.Config
	conf.MustLoad(confPath, &c)
	svc.InitSvcCtx(c)
}
