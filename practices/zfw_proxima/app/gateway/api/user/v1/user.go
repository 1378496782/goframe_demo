package api

import "github.com/gogf/gf/v2/frame/g"

type LoginReq struct {
	g.Meta   `path:"/users/login" method:"post" sm:"登录" tags:"用户"`
	Username string `json:"username" v:"required|length:3,12"`
	Password string `json:"password" v:"required|length:6,16"`
}

type LoginRes struct {
	Token string `json:"token" dc:"在需要鉴权的接口中header加入Authorization: token"`
}
