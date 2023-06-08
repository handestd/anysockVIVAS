package logic

import (
	"anysock/internal/model"
	"anysock/internal/svc"
	"anysock/internal/types"
	"anysock/pkg/httprequest"
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
)

type RandomUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRandomUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RandomUserLogic {
	return &RandomUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RandomUserLogic) RandomUser(req *types.UserRandomReq) (resp *types.UserRandomResp, err error) {
	body := httprequest.Get("https://randomuser.me/api/?nat=us")

	var randomUser model.RandomUser
	err = json.Unmarshal(body, &randomUser)
	if err != nil {
		return nil, err
	}

	return &types.UserRandomResp{
		Name:       randomUser.Results[0].Name.First + " " + randomUser.Results[0].Name.Last,
		Address:    strconv.Itoa(randomUser.Results[0].Location.Street.Number) + " " + randomUser.Results[0].Location.Street.Name,
		City:       randomUser.Results[0].Location.City,
		State:      randomUser.Results[0].Location.State,
		PostalCode: strconv.Itoa(randomUser.Results[0].Location.Postcode),
		Email:      randomUser.Results[0].Email,
		Password:   randomUser.Results[0].Login.Password,
		Phone:      randomUser.Results[0].Phone,
	}, nil
}
