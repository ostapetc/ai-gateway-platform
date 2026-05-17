package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	Name     string             `json:"name"`
	PostsRpc zrpc.RpcClientConf `json:"PostsRpc"`
	UsersRpc zrpc.RpcClientConf `json:"UsersRpc"`
}
