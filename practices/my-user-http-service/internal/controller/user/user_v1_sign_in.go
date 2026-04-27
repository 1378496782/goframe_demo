package user

import (
	"context"

	v1 "my-user-http-service/api/user/v1"
	"my-user-http-service/internal/service/user"
)

func (c *ControllerV1) SignIn(ctx context.Context, req *v1.SignInReq) (res *v1.SignInRes, err error) {
	err = c.userSvc.SignIn(ctx, user.SignInInput{
		Passport: req.Passport,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	return
}
