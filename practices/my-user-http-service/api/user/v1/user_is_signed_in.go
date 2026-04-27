package v1

import "github.com/gogf/gf/v2/frame/g"

type IsSignedInReq struct {
	g.Meta `path:"/user/is-signed-in" method:"post" tags:"UserService" summary:"Check if user is signed in"`
}

type IsSignedInRes struct {
	OK bool `dc:"True if current user is signed in; or else false"`
}
