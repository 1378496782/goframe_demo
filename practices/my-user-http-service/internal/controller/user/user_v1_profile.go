package user

import (
	"context"

	v1 "my-user-http-service/api/user/v1"

	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerV1) Profile(ctx context.Context, req *v1.ProfileReq) (res *v1.ProfileRes, err error) {
	user := c.userSvc.GetProfile(ctx)
	if user == nil {
		return nil, gerror.New("user not found")
	}
	res = &v1.ProfileRes{
		User: user,
	}
	return
}
