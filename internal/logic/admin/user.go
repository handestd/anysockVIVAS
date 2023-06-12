package admin

import (
	error_entity "anysock/internal/error"
	"anysock/internal/model"
	"anysock/internal/types"
	"anysock/pkg/cache"
	"context"
	"strconv"
)

func User(req *types.GetUserReq, l *MyUserLogic) (resp *types.GetUserResp, err error) {
	var id string

	if req.Id != "" {
		id = req.Id
	} else {
		userInfo := cache.GetMultiple("localUser")
		id = userInfo["UserId"]
	}

	idInt, myErr := strconv.ParseInt(id, 10, 64)

	if myErr == nil {
		var u *model.User
		u, err2 := l.svcCtx.UserModel.FindOne(context.Background(), idInt)

		if err2 != nil {
			// invalid user id error
			return &types.GetUserResp{}, error_entity.RecordNotFound.Error
		}
		//l.Logger.Info(u)

		return &types.GetUserResp{
			Username:  u.Username,
			Email:     u.Email,
			Balance:   u.Balance,
			UserAgent: l.ctx.Value("User-Agent").(string),
		}, nil
	} else {
		return &types.GetUserResp{}, error_entity.InValidUserID.Error
	}

	return
}
