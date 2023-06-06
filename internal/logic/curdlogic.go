package logic

import (
	"anysock/internal/model"
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"anysock/internal/svc"
	"anysock/internal/types"

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
	dsn := "root:@tcp(localhost:3306)/anysock?charset=utf8mb4&parseTime=True&loc=Local"
	conn := sqlx.NewMysql(dsn)

	var u model.User
	query := "select id, username from user where id=?"
	err = conn.QueryRowCtx(context.Background(), &u, query, 1)

	l.Logger.Info(u)
	return
}
