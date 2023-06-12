package auth

import (
	"anysock/pkg/cache"
	"context"
	"net/http"

	"anysock/internal/svc"
	"anysock/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogoutLogic) Logout(req *types.EmptyReq, w http.ResponseWriter) (resp *types.TextResp, err error) {
	// todo: add your logic here and delete this line
	sessionId := l.ctx.Value("data").(map[string]string)["session-id"]

	http.SetCookie(w, &http.Cookie{
		Name:  "session-id",
		Value: "",
	})

	//rediscache.DelKey()
	cache.DelKey(sessionId, l.ctx)
	l.Logger.Info(sessionId)

	l.Logger.Infov("---------------------------")
	l.Logger.Infov(sessionId)
	return
}
