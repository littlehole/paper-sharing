package logic

import (
	"context"
	"github.com/littlehole/paper-sharing/api/internal/svc"
	"github.com/littlehole/paper-sharing/api/internal/types"
	"github.com/littlehole/paper-sharing/rpc/user/userclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.UserClient.Register(l.ctx, &userclient.RegisterRequest{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		l.Logger.Errorf("user register err: %v", err)
		return nil, err
	}

	return &types.RegisterResponse{
		Username:  res.Username,
		CreatedAt: res.CreateAt,
		Message:   res.Message,
		JwtToken: types.JwtToken{
			AccessToken:  res.Jwt.AccessToken,
			AccessExpire: res.Jwt.AccessExpire,
			RefreshAfter: res.Jwt.RefreshAfter,
		},
	}, err
}
