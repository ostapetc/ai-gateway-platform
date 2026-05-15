package logic

import (
	"context"
	"fmt"
	"math/rand/v2"

	"github.com/brianvoe/gofakeit/v6"
	comments "github.com/ostapetc/ai-gateway-platform/services/comments/grpc/comments"
	"github.com/ostapetc/ai-gateway-platform/services/comments-bot-cronjob/internal/svc"
	"github.com/spf13/cobra"
)

func CreateComments(_ *cobra.Command, _ []string) error {
	ctx := context.Background()
	sc := svc.GetSvcCtx()

	req := &comments.CreateRequest{
		UserId: int64(rand.IntN(10) + 1),
		PostId: int64(rand.IntN(100) + 1),
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
