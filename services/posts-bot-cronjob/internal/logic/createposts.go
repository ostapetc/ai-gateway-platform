package logic

import (
	"context"
	"fmt"
	"math/rand/v2"

	"github.com/ostapetc/ai-gateway-platform/services/posts-bot-cronjob/internal/svc"
	posts "github.com/ostapetc/ai-gateway-platform/services/posts/grpc/posts"
	"github.com/spf13/cobra"
)

var (
	postTitles = []string{
		"Hello World",
		"Go is awesome",
		"Microservices rock",
		"AI Gateway update",
		"Random thoughts",
	}
	postBodies = []string{
		"This is a bot-generated post.",
		"Exploring the depths of distributed systems.",
		"Just another day in the life of a bot.",
		"Go routines are life.",
		"NATS JetStream makes async easy.",
	}
)

func CreatePosts(_ *cobra.Command, _ []string) error {
	ctx := context.Background()
	sc := svc.GetSvcCtx()

	req := &posts.AddRequest{
		UserId: uint64(rand.IntN(10) + 1),
		Title:  postTitles[rand.IntN(len(postTitles))],
		Body:   postBodies[rand.IntN(len(postBodies))],
	}

	resp, err := sc.PostsClient.Add(ctx, req)
	if err != nil {
		return fmt.Errorf("createPosts: %w", err)
	}

	fmt.Printf("created post id=%d\n", resp.Id)

	return nil
}
