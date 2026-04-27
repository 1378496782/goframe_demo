// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user

import (
	"context"

	"my-user-http-service/api/user/v1"
)

type IUserV1 interface {
	List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error)
	SignIn(ctx context.Context, req *v1.SignInReq) (res *v1.SignInRes, err error)
	SignUp(ctx context.Context, req *v1.SignUpReq) (res *v1.SignUpRes, err error)
}
