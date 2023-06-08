package logic

import (
	"anysock/internal/model"
	"anysock/pkg"
	"anysock/pkg/myredis"
	"context"
	"fmt"

	"anysock/internal/svc"
	"anysock/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req *types.GetUserReq) (resp *types.GetUserResp, err error) {
	// todo: add your logic here and delete this line

	UserId, myerr := myredis.GetSingleKey("UserId")

	UserIdInt, _ := pkg.InterfaceToInt64(UserId)

	if myerr == nil {
		var u *model.User
		u, err2 := l.svcCtx.UserModel.FindOne(context.Background(), UserIdInt)

		if err2 != nil {
			message := err2.Error()
			fmt.Println(message)
		}
		l.Logger.Info(u)

		return &types.GetUserResp{
			Username: u.Username,
			Email:    u.Email,
			Balance:  u.Balance,
		}, nil

	}

	return
}
