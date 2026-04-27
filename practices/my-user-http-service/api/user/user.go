// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user

import (
	"context"

	"my-user-http-service/api/user/v1"
)

type IUserV1 interface {
	IsSignedIn(ctx context.Context, req *v1.IsSignedInReq) (res *v1.IsSignedInRes, err error)
	List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error)
	Profile(ctx context.Context, req *v1.ProfileReq) (res *v1.ProfileRes, err error)
	SignIn(ctx context.Context, req *v1.SignInReq) (res *v1.SignInRes, err error)
	SignOut(ctx context.Context, req *v1.SignOutReq) (res *v1.SignOutRes, err error)
	SignUp(ctx context.Context, req *v1.SignUpReq) (res *v1.SignUpRes, err error)
}
