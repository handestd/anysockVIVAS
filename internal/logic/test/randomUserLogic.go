package test

import (
	"context"

	"anysock/internal/svc"
	"anysock/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RandomUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRandomUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RandomUserLogic {
	return &RandomUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RandomUserLogic) RandomUser(req *types.UserRandomReq) (resp *types.UserRandomResp, err error) {
	// todo: add your logic here and delete this line

	return
}
