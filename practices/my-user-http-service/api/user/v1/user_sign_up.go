package v1

import "github.com/gogf/gf/v2/frame/g"

type SignUpReq struct {
	g.Meta    `path:"/user/sign-up" method:"post" tags:"UserService" summary:"Sign up a new user account"`
	Passport  string `v:"required|length:6,16"  description:"User Passport"`
	Password  string `v:"required|length:6,16" description:"User Password"`
	Password2 string `v:"required|length:6,16|same:Password" description:"User Password Confirm"`
	Nickname  string `description:"User Nickname"`
}

type SignUpRes struct {
}
