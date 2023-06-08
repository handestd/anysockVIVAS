package logic

import (
	ErrorEntity "anysock/internal/error"
	"anysock/internal/model"
	"anysock/internal/svc"
	"anysock/internal/types"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
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
	// todo: add your logic here and delete this line
	// code check nil here
	// validate field here
	userExist, err := l.svcCtx.UserModel.FindByUsername(context.Background(), req.Username)

	if err != nil {
		// user no exist
	}

	if userExist == nil {
		password, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 14)

		u := &model.User{
			Username: req.Username,
			Password: string(password),
			Balance:  req.Balance,
			Email:    req.Email,
		}
		_, err = l.svcCtx.UserModel.Insert(l.ctx, u)
		//l.Logger.Info(u)

		if err != nil {
			return
		}

		return &types.RegisterResp{Name: u.Username}, nil
	} else {
		return &types.RegisterResp{ErrorMessage: ErrorEntity.UserExist.Message}, nil
	}
	//
}
