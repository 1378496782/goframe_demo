package user

import (
	"context"

	v1 "quick-demo/api/user/v1"
	"quick-demo/internal/dao"
	"quick-demo/internal/model/do"
)

func (c *ControllerV1) CreateUser(ctx context.Context, req *v1.CreateUserReq) (res *v1.CreateUserRes, err error) {
	res = &v1.CreateUserRes{}
	insertId, err := dao.User.Ctx(ctx).Data(do.User{
		Name:   req.Name,
		Status: v1.StatusOK,
		Age:    req.Age,
	}).InsertAndGetId()
	if err != nil {
		return nil, err
	}

	err = dao.User.Ctx(ctx).WherePri(insertId).Scan(&res.Entity)
	return
}
