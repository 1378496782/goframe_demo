package user

import (
	"context"

	v1 "my-user-http-service/api/user/v1"
)

func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {
	users, err := c.userSvc.GetList(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.ListRes{UserList: users}, nil
}
