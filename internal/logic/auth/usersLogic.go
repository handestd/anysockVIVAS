package auth

import (
	error_entity "anysock/internal/error"
	"anysock/internal/model"
	"context"
	"fmt"

	"anysock/internal/svc"
	"anysock/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UsersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UsersLogic {
	return &UsersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UsersLogic) Users(req *types.GetUsersReq) (resp *types.GetUsersResp, err error) {
	// sua lai cach dat ten bien
	var users []model.User

	// dat chung bien trong 1 scope
	users, err = l.svcCtx.UserModel.GetUsers(context.Background())

	if err != nil {
		message := err.Error()
		fmt.Println(message)
	}
	// tim hieu cach su dung logger
	//l.Logger.Info(u)

	var usersResp []types.GetUserResp
	//user2 := make([]types.GetUserResp,0) slice
	for _, v := range users {
		usersResp = append(usersResp, types.GetUserResp{Username: v.Username, Email: v.Email, Balance: v.Balance})
	}

	if len(usersResp) == 0 {
		return &types.GetUsersResp{}, error_entity.EmptyTable.Error
	}

	return &types.GetUsersResp{
		Users: usersResp,
	}, nil
}
