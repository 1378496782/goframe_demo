package v1

import (
	"my-user-http-service/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type ListReq struct {
	g.Meta `path:"/user/list" method:"get" tags:"UserService" summary:"Get the list of users"`
}
type ListRes struct {
	UserList []*entity.User `json:"userList" dc:"user list"`
}
