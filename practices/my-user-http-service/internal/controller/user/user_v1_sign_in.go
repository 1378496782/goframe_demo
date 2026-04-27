package user

import (
	"context"

	v1 "my-user-http-service/api/user/v1"
	usersvc "my-user-http-service/internal/service/user"
)

// SignIn signs in the user.
func (c *ControllerV1) SignIn(ctx context.Context, req *v1.SignInReq) (res *v1.SignInRes, err error) {
	err = c.userSvc.SignIn(ctx, usersvc.SignInInput{
		Passport: req.Passport,
		Password: req.Password,
	})
	return
}
