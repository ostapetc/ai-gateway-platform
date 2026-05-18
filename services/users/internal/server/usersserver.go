package server

import (
	"context"

	"github.com/ostapetc/ai-gateway-platform/services/users/grpc/users"
	"github.com/ostapetc/ai-gateway-platform/services/users/internal/logic"
	"github.com/ostapetc/ai-gateway-platform/services/users/internal/svc"
	"github.com/ostapetc/ai-gateway-platform/services/users/internal/types"
)

type UsersServer struct {
	svcCtx *svc.ServiceContext
	users.UnimplementedUsersServer
}

func NewUsersServer(svcCtx *svc.ServiceContext) *UsersServer {
	return &UsersServer{svcCtx: svcCtx}
}

func (s *UsersServer) Create(ctx context.Context, in *users.CreateRequest) (*users.CreateResponse, error) {
	l := logic.NewCreateUserLogic(ctx, s.svcCtx)
	resp, err := l.CreateUser(&types.CreateUserRequest{
		Username: in.Username,
		Email:    in.Email,
		Password: in.Password,
	})
	if err != nil {
		return nil, err
	}
	return &users.CreateResponse{Id: resp.ID}, nil
}

func (s *UsersServer) Get(ctx context.Context, in *users.GetRequest) (*users.GetResponse, error) {
	u, ok := s.svcCtx.UserStore.Get(in.Id)
	if !ok {
		return &users.GetResponse{}, nil
	}
	return &users.GetResponse{
		User: &users.User{
			Id:       u.ID,
			Username: u.Username,
			Email:    u.Email,
		},
	}, nil
}

func (s *UsersServer) GetRandom(ctx context.Context, in *users.GetRandomRequest) (*users.GetRandomResponse, error) {
	l := logic.NewGetRandomUserLogic(ctx, s.svcCtx)
	resp, err := l.GetRandomUser()
	if err != nil {
		return nil, err
	}
	return &users.GetRandomResponse{
		User: &users.User{
			Id:       resp.User.ID,
			Username: resp.User.Username,
			Email:    resp.User.Email,
		},
	}, nil
}

func (s *UsersServer) List(ctx context.Context, in *users.ListRequest) (*users.ListResponse, error) {
	l := logic.NewListUsersLogic(ctx, s.svcCtx)
	resp, err := l.ListUsers()
	if err != nil {
		return nil, err
	}

	result := make([]*users.User, len(resp.Users))
	for i, u := range resp.Users {
		result[i] = &users.User{
			Id:       u.ID,
			Username: u.Username,
			Email:    u.Email,
		}
	}
	return &users.ListResponse{Users: result, Total: uint64(resp.Total)}, nil
}
