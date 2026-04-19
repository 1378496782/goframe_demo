package users

import (
	"context"

	v1 "star/api/users/v1"
)

func (c *ControllerV1) DeleteUser(ctx context.Context, req *v1.DeleteUserReq) (res *v1.DeleteUserRes, err error) {
	err = c.users.DeleteUser(ctx, req.Id)
	return &v1.DeleteUserRes{}, err
}
