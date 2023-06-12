// Code generated by goctl. DO NOT EDIT.
package types

type EmptyReq struct {
	Message string `json:"message"`
}

type TextResp struct {
	Message string `json:"message"`
}

type UserRandomReq struct {
}

type UserRandomResp struct {
	Name       string `json:"name"`
	Address    string `json:"address"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"postal_code"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
}

type GetUserReq struct {
	Id string `path:"id,optional"`
}

type GetUserResp struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Balance   int64  `json:"balance"`
	UserAgent string `json:"useragent"`
}

type GetUsersReq struct {
}

type GetUsersResp struct {
	Users []GetUserResp `json:"users"`
}

type RegisterReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type RegisterResp struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	ErrorMessage string `json:"errorMessage"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResp struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	AccessToken  string `json:"access_token"`
	AccessExpire int64  `json:"access_expire"`
	RefreshAfter int64  `json:"refresh_after"`
	ErrorMessage string `json:"errorMessage"`
}
