// Code generated by goctl. DO NOT EDIT.
package types

type RegisterReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterResp struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Token    string `json:"token"`
	ExpireAt string `json:"expireAt"`
}

type PayloadReq struct {
	Message string `json:"message"`
}

type TextResp struct {
	Message string `json:"message"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResp struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Token    string `json:"token"`
	ExpireAt string `json:"expireAt"`
}