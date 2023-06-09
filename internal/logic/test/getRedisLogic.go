package test

import (
	"context"

	"anysock/internal/svc"
	"anysock/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRedisLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRedisLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRedisLogic {
	return &GetRedisLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRedisLogic) GetRedis(req *types.PayloadReq) (resp *types.TextResp, err error) {
	// todo: add your logic here and delete this line

	return
}
