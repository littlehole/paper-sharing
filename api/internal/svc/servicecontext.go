package svc

import (
	"github.com/littlehole/paper-sharing/api/internal/config"
	"github.com/littlehole/paper-sharing/rpc/user/userclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	UserClient userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		UserClient: userclient.NewUser(zrpc.MustNewClient(c.User)),
	}
}
