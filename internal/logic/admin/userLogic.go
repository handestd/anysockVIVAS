package admin

import (
	error_entity "anysock/internal/error"
	"anysock/internal/model"
	"anysock/internal/svc"
	"anysock/internal/types"
	"anysock/pkg/cache"
	"context"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogic) User(req *types.GetUserReq) (resp *types.GetUserResp, err error) {
	UserInfo := cache.GetMultiple("localUser")
	UserIdInt, myerr := strconv.ParseInt(UserInfo["UserId"], 10, 64)
	//UserIdInt, _ := pkg.InterfaceToInt64(UserId)

	if myerr == nil {
		var u *model.User
		u, err2 := l.svcCtx.UserModel.FindOne(context.Background(), UserIdInt)

		if err2 != nil {
			// invalid user id error
			return &types.GetUserResp{}, error_entity.RecordNotFound.Error
		}
		//l.Logger.Info(u)

		return &types.GetUserResp{
			Username: u.Username,
			Email:    u.Email,
			Balance:  u.Balance,
		}, nil
	} else {
		return &types.GetUserResp{}, error_entity.InValidUserID.Error
	}

	return
}
