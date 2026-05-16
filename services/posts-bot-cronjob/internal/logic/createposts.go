package logic

import (
	"context"
	"fmt"
	"math/rand/v2"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ostapetc/ai-gateway-platform/services/posts-bot-cronjob/internal/svc"
	posts "github.com/ostapetc/ai-gateway-platform/services/posts/grpc/posts"
	"github.com/spf13/cobra"
)

func CreatePosts(_ *cobra.Command, _ []string) error {
	ctx := context.Background()
	sc := svc.GetSvcCtx()

	for range 3 {
		err := createPost(ctx, sc)
		if err != nil {
			return err
		}
	}

	return nil
}

func createPost(ctx context.Context, sc *svc.ServiceContext) error {
	req := &posts.AddRequest{
		UserId: uint64(rand.IntN(10) + 1),
		Title:  generateTitle(),
		Body:   generateBody(),
	}

	resp, err := sc.PostsClient.Add(ctx, req)
	if err != nil {
		return fmt.Errorf("createPosts: %w", err)
	}

	fmt.Printf("created post id=%d\n", resp.Id)

	return nil
}

func generateTitle() string {
	return gofakeit.Sentence(rand.IntN(4) + 5)
}

func generateBody() string {
	return gofakeit.Paragraph(1, rand.IntN(3)+2, rand.IntN(5)+5, " ")
}
