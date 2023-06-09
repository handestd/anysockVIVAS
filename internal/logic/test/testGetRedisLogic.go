package test

import (
	"context"

	"anysock/internal/svc"
	"anysock/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestGetRedisLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTestGetRedisLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestGetRedisLogic {
	return &TestGetRedisLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TestGetRedisLogic) TestGetRedis(req *types.PayloadReq) (resp *types.TextResp, err error) {
	// todo: add your logic here and delete this line

	return
}
