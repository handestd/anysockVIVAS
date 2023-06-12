package user

import (
	error_entity "anysock/internal/error"
	"anysock/internal/model"
	"context"
	"golang.org/x/crypto/bcrypt"

	"anysock/internal/svc"
	"anysock/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	userExist, err := l.svcCtx.UserModel.FindByUsername(context.Background(), req.Username)

	if err != nil {
		// user no exist
	}

	if userExist == nil {
		password, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 14)

		u := &model.User{
			Username: req.Username,
			Password: string(password),
			Email:    req.Email,
		}
		_, err = l.svcCtx.UserModel.Insert(l.ctx, u)
		//l.Logger.Info(u)

		if err != nil {
			return
		}

		return &types.RegisterResp{Name: u.Username}, nil
	} else {
		return &types.RegisterResp{ErrorMessage: error_entity.UserExist.Message}, nil
	}
	return
}
