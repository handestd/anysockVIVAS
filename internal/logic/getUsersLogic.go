package logic

import (
	ErrorEntity "anysock/internal/error"
	"anysock/internal/model"
	"anysock/internal/svc"
	"anysock/internal/types"
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUsersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUsersLogic {
	return &GetUsersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUsersLogic) GetUsers(req *types.GetUsersReq) (resp *types.GetUsersResp, err error) {

	var u []model.User
	u, err2 := l.svcCtx.UserModel.GetUsers(context.Background())

	if err2 != nil {
		message := err2.Error()
		fmt.Println(message)
	}
	//l.Logger.Info(u)

	var users []types.GetUserResp
	//user2 := make([]types.GetUserResp,0) slice
	for _, v := range u {
		users = append(users, types.GetUserResp{Username: v.Username, Email: v.Email, Balance: v.Balance})
	}

	if len(users) == 0 {
		return &types.GetUsersResp{}, ErrorEntity.EmptyTable.Error
	}

	return &types.GetUsersResp{
		Users: users,
	}, nil

	//return
}