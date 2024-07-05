package svc

import (
	"github.com/littlehole/paper-sharing/internal/db"
	"github.com/littlehole/paper-sharing/internal/rpc/user/internal/config"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DB:     db.NewDB(),
	}
}
