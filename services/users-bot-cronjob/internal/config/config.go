package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	Name     string             `json:"name"`
	UsersRpc zrpc.RpcClientConf `json:"UsersRpc"`
}
