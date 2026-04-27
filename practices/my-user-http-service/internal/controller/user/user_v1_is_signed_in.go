package user

import (
	"context"

	v1 "my-user-http-service/api/user/v1"
)

func (c *ControllerV1) IsSignedIn(ctx context.Context, req *v1.IsSignedInReq) (res *v1.IsSignedInRes, err error) {
	return &v1.IsSignedInRes{
		OK: c.userSvc.IsSignedIn(ctx),
	}, nil
}
