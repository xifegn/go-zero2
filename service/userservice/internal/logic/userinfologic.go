package logic

import (
	"context"

	"go-zero2/service/userservice/internal/svc"
	"go-zero2/service/userservice/pb/userservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *userservice.UserInfoRequest) (*userservice.UserInfoResponse, error) {
	// todo: add your logic here and delete this line

	return &userservice.UserInfoResponse{}, nil
}
