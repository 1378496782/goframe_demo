package v1

import (
	"star/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type RegisterReq struct {
	g.Meta   `path:"/users/register" method:"post" sm:"注册" tags:"用户"`
	Username string `json:"username" v:"required|length:3,12" dc:"用户名"`
	Password string `json:"password" v:"required|length:6,16" dc:"密码"`
	Email    string `json:"email" v:"required|email" dc:"邮箱"`
}

type RegisterRes struct {
}

type UserListReq struct {
	g.Meta `path:"/users/list" method:"get" sm:"获取用户列表" tags:"用户"`
}

type UserListRes struct {
	UserList []*entity.Users `json:"userList" dc:"用户列表"`
}

type DeleteUserReq struct {
	g.Meta `path:"/users/{id}" method:"delete" sm:"删除用户" tags:"用户"`
	Id     int64 `json:"id" v:"required" dc:"user id"`
}

type DeleteUserRes struct {
}

type GetUserReq struct {
	g.Meta `path:"/users/{id}" tags:"用户" method:"GET" summary:"获取用户信息"`
	Id     int64 `json:"id" v:"required" dc:"user id"`
}

type GetUserRes struct {
	User *entity.Users `json:"user" dc:"user"`
}
