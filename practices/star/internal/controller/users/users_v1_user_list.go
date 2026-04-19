package users

import (
	"context"

	v1 "star/api/users/v1"
)

func (c *ControllerV1) UserList(ctx context.Context, req *v1.UserListReq) (res *v1.UserListRes, err error) {
	res = &v1.UserListRes{}
	res.UserList, err = c.users.GetUserList(ctx)
	return
}
