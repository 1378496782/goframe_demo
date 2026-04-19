package users

import (
	"context"

	v1 "star/api/users/v1"
)

func (c *ControllerV1) GetUser(ctx context.Context, req *v1.GetUserReq) (res *v1.GetUserRes, err error) {
	res = &v1.GetUserRes{}
	res.User, err = c.users.GetUser(ctx, req.Id)
	return
}
