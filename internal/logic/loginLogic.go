package logic

import (
	ErrorEntity "anysock/internal/error"
	"anysock/internal/svc"
	"anysock/internal/types"
	"anysock/pkg/myredis"
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// todo: add your logic here and delete this line
	userExist, err := l.svcCtx.UserModel.FindByUsername(context.Background(), req.Username)

	if err != nil {
		return
	}

	if userExist != nil {
		err := bcrypt.CompareHashAndPassword([]byte(userExist.Password), []byte(req.Password))
		if err != nil {
			return &types.LoginResp{ErrorMessage: ErrorEntity.IncorrectPassword.Message}, nil
		}
		accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire

		now := time.Now().Unix()
		accessToken, err := l.GenToken(now, l.svcCtx.Config.JwtAuth.AccessSecret, nil, accessExpire)

		myredis.SetSinglekey("UserId", userExist.Id)
		myredis.SetSinglekey("UserAccessToken", accessToken)
		myredis.SetSinglekey("AccessExpire", now+accessExpire)
		myredis.SetSinglekey("RefreshAfter", now+accessExpire/2)

		return &types.LoginResp{
			Id:           userExist.Id,
			Name:         req.Username,
			AccessToken:  accessToken,
			AccessExpire: now + accessExpire,
			RefreshAfter: now + accessExpire/2,
		}, nil
	}

	return
}

func (l *LoginLogic) GenToken(iat int64, secretKey string, payloads map[string]interface{}, seconds int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	for k, v := range payloads {
		claims[k] = v
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	return token.SignedString([]byte(secretKey))
}
