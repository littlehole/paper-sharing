package logic

import (
	"context"

	"github.com/littlehole/paper-sharing/api/internal/svc"
	"github.com/littlehole/paper-sharing/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequst) (resp *types.LoginResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
