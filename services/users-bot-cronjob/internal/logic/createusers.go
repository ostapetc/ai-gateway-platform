package logic

import (
	"context"
	"fmt"
	"math/rand/v2"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ostapetc/ai-gateway-platform/services/users-bot-cronjob/internal/svc"
	users "github.com/ostapetc/ai-gateway-platform/services/users/grpc/users"
	"github.com/spf13/cobra"
)

func CreateUsers(_ *cobra.Command, _ []string) error {
	ctx := context.Background()
	sc := svc.GetSvcCtx()

	req := &users.CreateRequest{
		Username: gofakeit.Username(),
		Email:    gofakeit.Email(),
		Password: generatePassword(),
	}

	resp, err := sc.UsersClient.Create(ctx, req)
	if err != nil {
		return fmt.Errorf("createUsers: %w", err)
	}

	fmt.Printf("created user id=%d\n", resp.Id)

	return nil
}

func generatePassword() string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 16+rand.IntN(8))
	for i := range b {
		b[i] = chars[rand.IntN(len(chars))]
	}
	return string(b)
}
