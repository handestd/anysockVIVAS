package test

import (
	"context"

	"anysock/internal/svc"
	"anysock/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetRedisLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetRedisLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetRedisLogic {
	return &SetRedisLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetRedisLogic) SetRedis(req *types.PayloadReq) (resp *types.TextResp, err error) {
	// todo: add your logic here and delete this line

	return
}
