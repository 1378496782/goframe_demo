package v1

import (
	"quick-demo/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type GetUserReq struct {
	g.Meta `path:"/user/{id}" tags:"User" method:"GET" summary:"获取用户信息"`
	Id     int64 `json:"id" v:"required" dc:"user id"`
}

type GetUserRes struct {
	Entity *entity.User `json:"user" dc:"user"`
}

type GetUserListReq struct {
	g.Meta `path:"/user/list" tags:"User" method:"GET" summary:"获取用户列表"`
}

type GetUserListRes struct {
	UserList []*entity.User `json:"userList" dc:"user list"`
}

type CreateUserReq struct {
	g.Meta `path:"/user" tags:"User" method:"POST" summary:"创建用户"`
	Name   string `v:"required|length:3,10" dc:"user name"`
	Age    uint   `v:"required|between:18,200" dc:"user age"`
}

type CreateUserRes struct {
	Entity *entity.User `json:"user" dc:"user"`
}

type Status int

const (
	StatusOK       Status = 0 // User is OK.
	StatusDisabled Status = 1 // User is disabled.
)
