package user

import (
	"context"

	v1 "my-user-http-service/api/user/v1"
)

func (c *ControllerV1) SignOut(ctx context.Context, req *v1.SignOutReq) (res *v1.SignOutRes, err error) {
	err = c.userSvc.SignOut(ctx)
	return
}
