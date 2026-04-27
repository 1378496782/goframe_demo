package v1

import "github.com/gogf/gf/v2/frame/g"

type SignInReq struct {
	g.Meta   `path:"/user/sign-in" method:"post" tags:"UserService" summary:"Sign in with exist account"`
	Passport string `v:"required" dc:"User Passport"`
	Password string `v:"required" dc:"User Password"`
}

type SignInRes struct {
}
