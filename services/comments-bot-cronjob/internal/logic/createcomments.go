package logic

import (
	"context"
	"fmt"
	"math/rand/v2"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ostapetc/ai-gateway-platform/services/comments-bot-cronjob/internal/svc"
	comments "github.com/ostapetc/ai-gateway-platform/services/comments/grpc/comments"
	"github.com/ostapetc/ai-gateway-platform/services/posts/grpc/posts"
	"github.com/spf13/cobra"
)

func CreateComments(_ *cobra.Command, _ []string) error {
	ctx := context.Background()
	sc := svc.GetSvcCtx()

	postsResp, err := sc.PostsClient.GetRandom(ctx, &posts.GetRandomRequest{})
	if err != nil {
		return fmt.Errorf("get random post error: %w", err)
	}

	post := postsResp.Post

	if post == nil {
		return fmt.Errorf("no random posts found")
	}

	req := &comments.CreateRequest{
		UserId: uint64(rand.IntN(10) + 1),
		PostId: post.Id,
		Body:   generateBody(),
	}

	resp, err := sc.CommentsClient.Create(ctx, req)
	if err != nil {
		return fmt.Errorf("createComments: %w", err)
	}

	fmt.Printf("created comment id=%d\n", resp.Id)

	return nil
}

func generateBody() string {
	return gofakeit.Paragraph(1, rand.IntN(3)+1, rand.IntN(5)+5, " ")
}
