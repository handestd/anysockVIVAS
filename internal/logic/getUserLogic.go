package logic

import (
	ErrorEntity "anysock/internal/error"
	"anysock/internal/model"
	"anysock/internal/svc"
	"anysock/internal/types"
	"anysock/pkg/myredis"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
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

	//UserId, myerr := myredis.GetSingleKey("UserId")

	UserInfo := myredis.GetMultiple("localUser")
	UserIdInt, myerr := strconv.ParseInt(UserInfo["UserId"], 10, 64)
	//UserIdInt, _ := pkg.InterfaceToInt64(UserId)

	if myerr == nil {
		var u *model.User
		u, err2 := l.svcCtx.UserModel.FindOne(context.Background(), UserIdInt)

		if err2 != nil {
			// invalid user id error
			return &types.GetUserResp{}, ErrorEntity.RecordNotFound.Error
		}
		//l.Logger.Info(u)

		return &types.GetUserResp{
			Username: u.Username,
			Email:    u.Email,
			Balance:  u.Balance,
		}, nil
	} else {
		return &types.GetUserResp{}, ErrorEntity.InValidUserID.Error
	}
}
