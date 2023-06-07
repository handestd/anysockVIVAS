package logic

import (
	"anysock/internal/svc"
	"anysock/internal/types"
	"anysock/pkg/myredis"
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestSetRedisLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTestSetRedisLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestSetRedisLogic {
	return &TestSetRedisLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TestSetRedisLogic) TestSetRedis(req *types.PayloadReq) (resp *types.TextResp, err error) {
	// todo: add your logic here and delete this line
	ctx := context.Background()

	err2 := myredis.Client.Set(ctx, "foo", "bar", 0).Err()
	if err2 != nil {
		panic(err2)
	}

	val, err2 := myredis.Client.Get(ctx, "foo").Result()
	if err2 != nil {
		panic(err2)
	}
	fmt.Println("foo", val)
	return
}
