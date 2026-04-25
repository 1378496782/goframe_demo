// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user

import (
	"context"

	v1 "zfw_proxima/app/gateway/api/user/v1"
)

type IUserV1 interface {
	Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error)
}
