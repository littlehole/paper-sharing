package logic

import (
	"context"
	"errors"
	database "github.com/littlehole/paper-sharing/internal/db"
	"github.com/littlehole/paper-sharing/internal/db/models"
	rpcerrors "github.com/littlehole/paper-sharing/internal/errors"
	"github.com/littlehole/paper-sharing/internal/rpc/user/internal/svc"
	"github.com/littlehole/paper-sharing/internal/rpc/user/user"
	"gorm.io/gorm"
	"k8s.io/klog/v2"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

var db = database.NewDB()

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
	// select user
	var userModel *models.UserModel
	var labModel *models.LabModel
	err := db.Model(models.UserModel{}).
		Where("lab_name = ? AND grade = ? AND name = ? AND username = ?",
			in.LabName, in.Grade, in.Name, in.Username).First(&userModel).Error
	if err == nil {
		klog.Error("register user error: user already exists")
		return nil, errors.New(rpcerrors.ErrUserExists)
	}

	userModel, labModel = converRegisterRequestToModel(in)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		if err := registerLabAndUser(userModel, labModel); err != nil {
			l.Logger.Errorf("register lab error: %v", err)
			return nil, err
		}
	} else {
		// unknown error
		l.Logger.Errorf("register user error: %v", err)
		return nil, err
	}

	return &user.RegisterResponse{
		Username: in.Username,
		CreateAt: time.Now().Format("2006-01-02 15:04:05"),
		Message:  "register successfully",
		Jwt:      &user.JwtToken{},
	}, nil
}

func registerLabAndUser(userModel *models.UserModel, labModel *models.LabModel) error {
	// select lab
	lab := &models.LabModel{}
	err := db.Where("lab_name = ?", userModel.LabName).First(lab).Error
	tx := db.Begin()
	if err == nil {
		// found in lab
		// create user
		// update lab
		if err := tx.Create(&userModel).Error; err != nil {
			tx.Rollback()
			return err
		}
		lab.UserList = append(lab.UserList, userModel.Username)
		if err := tx.Updates(&labModel).Error; err != nil {
			tx.Rollback()
			return err
		}
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		// not found in lab
		// create lab and user
		if err := tx.Create(&userModel).Error; err != nil {
			tx.Rollback()
			return err
		}
		if err := tx.Create(labModel).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func converRegisterRequestToModel(req *user.RegisterRequest) (*models.UserModel, *models.LabModel) {
	return &models.UserModel{
			LabName:   req.LabName,
			Grade:     req.Grade,
			Name:      req.Name,
			Username:  req.Username,
			Password:  req.Password,
			ShareList: []string{},
			PaperList: []string{},
		},
		&models.LabModel{
			LabName:  req.LabName,
			LabPass:  req.LabPass,
			UserList: []string{req.Username},
		}
}

func convertModelToRegisterResponse(u *models.UserModel) *user.RegisterResponse {
	return &user.RegisterResponse{
		Username: u.Username,
		CreateAt: time.Now().Format("2006-01-02 15:04:05"),
		Message:  "register successfully",
		Jwt:      &user.JwtToken{},
	}
}
