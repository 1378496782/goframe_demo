package logic

import (
	"context"

	v1 "my-user-http-service/api/user/v1"
	"my-user-http-service/internal/dao"
)

func GetList(ctx context.Context) (res *v1.ListRes, err error) {
	res = &v1.ListRes{}
	err = dao.User.Ctx(ctx).Scan(&res.UserList)
	return
}
