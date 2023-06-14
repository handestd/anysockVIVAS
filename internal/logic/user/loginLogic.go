package user

import (
	error_entity "anysock/internal/error"
	"anysock/internal/model"
	"anysock/pkg/cache"
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"

	"anysock/internal/svc"
	"anysock/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *LoginLogic) Login(req *types.LoginReq, w http.ResponseWriter) (resp *types.LoginResp, err error) {
	userExist, err := l.svcCtx.UserModel.FindByUsername(context.Background(), req.Username)

	if err != nil {
		return
	}

	if userExist != nil {
		err := bcrypt.CompareHashAndPassword([]byte(userExist.Password), []byte(req.Password))
		if err != nil {
			return &types.LoginResp{ErrorMessage: error_entity.IncorrectPassword.Message}, nil
		}
		accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire

		now := time.Now().Unix()
		accessToken, err := l.GenToken(now, l.svcCtx.Config.JwtAuth.AccessSecret, nil, accessExpire)

		//cache.SetSinglekey("UserId", userExist.Id)
		//cache.SetSinglekey("UserAccessToken", accessToken)
		//cache.SetSinglekey("AccessExpire", now+accessExpire)
		//cache.SetSinglekey("RefreshAfter", now+accessExpire/2)

		//session := map[string]interface{}{
		//	"UserId":          userExist.Id,
		//	"UserAccessToken": accessToken,
		//	"AccessExpire":    now + accessExpire,
		//	"RefreshAfter":    now + accessExpire/2,
		//}

		//cache.SetMultiple(session, "localUser")
		l.CreateUserSession(w, *userExist)

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

func (l *LoginLogic) CreateUserSession(w http.ResponseWriter, user model.User) error {
	sessionID := uuid.New().String()
	var dataUser map[string]interface{}
	ctx := context.Background()

	http.SetCookie(w, &http.Cookie{
		Name:   "session-id",
		Value:  sessionID,
		MaxAge: 3600,
	})
	// convert struct to map
	dataUser = map[string]interface{}{
		"id":       user.Id,
		"username": user.Username,
		"auth":     true,
	}

	cache.SetMultiple(ctx, dataUser, sessionID)
	cache.Client.Expire(ctx, sessionID, time.Hour)

	return nil
}
