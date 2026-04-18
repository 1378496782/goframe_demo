package user

import (
	"context"
	"quick-demo/internal/dao"

	v1 "quick-demo/api/user/v1"
)

func (c *ControllerV1) GetUser(ctx context.Context, req *v1.GetUserReq) (res *v1.GetUserRes, err error) {
	res = &v1.GetUserRes{}
	err = dao.User.Ctx(ctx).WherePri(req.Id).Scan(&res.Entity)
	return res, err
}
