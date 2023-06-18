package logic

import (
	"context"

	"go-zero2/service/userservice/internal/svc"
	"go-zero2/service/userservice/pb/userservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *userservice.LoginRequest) (*userservice.LoginResponse, error) {
	// todo: add your logic here and delete this line

	return &userservice.LoginResponse{}, nil
}
