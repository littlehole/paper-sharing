package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/littlehole/paper-sharing/internal/db"
	"github.com/littlehole/paper-sharing/internal/db/models"
	"github.com/littlehole/paper-sharing/internal/rpc/user/internal/svc"
	"github.com/littlehole/paper-sharing/internal/rpc/user/user"
	"gorm.io/gorm"
	"k8s.io/klog/v2"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.RegisterResponse, error) {
	// todo: add your logic here and delete this line
	db := db.NewDB()
	var userModel models.UserModel
	err := db.Model(&models.UserModel{}).
		Where("lab_name = ? AND grade = ? AND name = ?", in.LabName, in.Grade, in.Name).First(&userModel).Error
	if err == nil {
		klog.Error("register user error: user already exists")
		return &user.RegisterResponse{
			Username: fmt.Sprintf("%s/%s/%s", in.LabName, in.Grade, in.Name),
			CreateAt: time.Now().Format("2006-01-02 15:04:05"),
			Message:  "fail to register user, reason: user already exists",
			Jwt:      &user.JwtToken{},
		}, nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// create
		err := db.Model(&models.UserModel{}).Create(&models.UserModel{
			UserInfo: models.UserInfo{
				LabName: in.LabName,
				Grade:   in.Grade,
				Name:    in.Name,
			},
			Password:  in.Password,
			ShareList: []string{},
			PaperList: []string{},
		}).Error
		if err != nil {
			klog.Errorf("create user data error: %v", err)
			return &user.RegisterResponse{
				Username: fmt.Sprintf("%s/%s/%s", in.LabName, in.Grade, in.Name),
				CreateAt: time.Now().Format("2006-01-02 15:04:05"),
				Message:  "fail to register user, reason: create user data error: " + err.Error(),
				Jwt:      &user.JwtToken{},
			}, nil
		}
		return &user.RegisterResponse{
			Username: fmt.Sprintf("%s/%s/%s", in.LabName, in.Grade, in.Name),
			CreateAt: time.Now().Format("2006-01-02 15:04:05"),
			Message:  "create successfully",
			Jwt:      &user.JwtToken{},
		}, nil
	}
	return &user.RegisterResponse{
		Username: fmt.Sprintf("%s/%s/%s", in.LabName, in.Grade, in.Name),
		CreateAt: time.Now().Format("2006-01-02 15:04:05"),
		Message:  "fail to register user, reason: " + err.Error(),
		Jwt:      &user.JwtToken{},
	}, err
}
