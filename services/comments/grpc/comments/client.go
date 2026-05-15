package comments

import (
	"context"

	"github.com/zeromicro/go-zero/zrpc"
)

type (
	Client interface {
		Create(ctx context.Context, in *CreateRequest) (*CreateResponse, error)
		List(ctx context.Context, in *ListRequest) (*ListResponse, error)
	}

	defaultClient struct {
		cli zrpc.Client
	}
)

func NewClient(cli zrpc.Client) Client {
	return &defaultClient{cli}
}

func (m *defaultClient) Create(ctx context.Context, in *CreateRequest) (*CreateResponse, error) {
	return NewCommentsClient(m.cli.Conn()).Create(ctx, in)
}

func (m *defaultClient) List(ctx context.Context, in *ListRequest) (*ListResponse, error) {
	return NewCommentsClient(m.cli.Conn()).List(ctx, in)
}
