package user

import (
	"context"

	account "zfw_proxima/app/user/api/account/v1"

	v1 "zfw_proxima/app/gateway/api/user/v1"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	user, err := c.AccountClient.UserLogin(ctx, &account.UserLoginReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	return &v1.LoginRes{
		Token: user.GetToken(),
	}, nil
}
