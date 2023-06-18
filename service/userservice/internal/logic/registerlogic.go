package logic

import (
	"context"
	"fmt"
	"go-zero2/internal/model/usermodel"
	"go-zero2/pkg/cryptx"
	"go-zero2/service/userservice/internal/svc"
	"go-zero2/service/userservice/pb/userservice"
	"go-zero2/service/userservice/user"
	"google.golang.org/grpc/status"

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

func (l *RegisterLogic) Register(in *userservice.RegisterRequest) (*userservice.RegisterResponse, error) {
	resp, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)
	fmt.Println(resp, err)
	if err == nil {
		return nil, status.Error(100, "User already exists")
	}
	if err == usermodel.ErrNotFound {
		newUser := usermodel.User{
			Name:     in.Name,
			Gender:   in.Gender,
			Mobile:   in.Mobile,
			Password: cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password),
		}
		res, err := l.svcCtx.UserModel.Insert(l.ctx, &newUser)
		if err != nil {
			return nil, status.Error(500, err.Error())
		}
		newUser.Id, err = res.LastInsertId()
		if err != nil {
			return nil, status.Error(500, err.Error())
		}
		return &user.RegisterResponse{
			Id:     newUser.Id,
			Name:   newUser.Name,
			Gender: newUser.Gender,
			Mobile: newUser.Mobile,
		}, nil
	}
	return nil, status.Error(100, err.Error())
}
