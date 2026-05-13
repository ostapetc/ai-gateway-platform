package posts

import (
	"context"

	"github.com/zeromicro/go-zero/zrpc"
)

type (
	Client interface {
		Add(ctx context.Context, in *AddRequest) (*AddResponse, error)
		List(ctx context.Context, in *ListRequest) (*ListResponse, error)
	}

	defaultClient struct {
		cli zrpc.Client
	}
)

func NewClient(cli zrpc.Client) Client {
	return &defaultClient{cli}
}

func (m *defaultClient) Add(ctx context.Context, in *AddRequest) (*AddResponse, error) {
	return NewPostsClient(m.cli.Conn()).Add(ctx, in)
}

func (m *defaultClient) List(ctx context.Context, in *ListRequest) (*ListResponse, error) {
	return NewPostsClient(m.cli.Conn()).List(ctx, in)
}
