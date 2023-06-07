package logic

import (
	"context"

	"anysock/internal/svc"
	"anysock/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShowAllLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShowAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShowAllLogic {
	return &ShowAllLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShowAllLogic) ShowAll(req *types.PayloadReq) (resp *types.TextResp, err error) {
	// todo: add your logic here and delete this line

	return
}
