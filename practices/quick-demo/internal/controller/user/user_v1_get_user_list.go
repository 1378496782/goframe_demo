package user

import (
	"context"

	v1 "quick-demo/api/user/v1"
	"quick-demo/internal/dao"
)

func (c *ControllerV1) GetUserList(ctx context.Context, req *v1.GetUserListReq) (res *v1.GetUserListRes, err error) {
	res = &v1.GetUserListRes{}
	err = dao.User.Ctx(ctx).Scan(&res.UserList)
	return
}
