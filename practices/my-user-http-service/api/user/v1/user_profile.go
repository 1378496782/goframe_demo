package v1

import (
	"my-user-http-service/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type ProfileReq struct {
	g.Meta `path:"/user/profile" method:"get" tags:"UserService" summary:"Get user profile"`
}

type ProfileRes struct {
	*entity.User
}
