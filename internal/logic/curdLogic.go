package logic

import (
	"anysock/internal/model"
	"anysock/internal/svc"
	"anysock/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type CurdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCurdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CurdLogic {
	return &CurdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CurdLogic) Curd(req *types.PayloadReq) (resp *types.TextResp, err error) {
	// todo: add your logic here and delete this line
	var u *model.User
	u, _ = l.svcCtx.UserModel.FindOne(l.ctx, 1)
	l.Logger.Info(u)
	return &types.TextResp{Message: "success"}, nil
}
