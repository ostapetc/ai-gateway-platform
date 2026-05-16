package users

import (
	"context"

	"github.com/zeromicro/go-zero/zrpc"
)

type (
	Client interface {
		Create(ctx context.Context, in *CreateRequest) (*CreateResponse, error)
		Get(ctx context.Context, in *GetRequest) (*GetResponse, error)
	}

	defaultClient struct {
		cli zrpc.Client
	}
)

func NewClient(cli zrpc.Client) Client {
	return &defaultClient{cli}
}

func (m *defaultClient) Create(ctx context.Context, in *CreateRequest) (*CreateResponse, error) {
	return NewUsersClient(m.cli.Conn()).Create(ctx, in)
}

func (m *defaultClient) Get(ctx context.Context, in *GetRequest) (*GetResponse, error) {
	return NewUsersClient(m.cli.Conn()).Get(ctx, in)
}
