package logic

import (
	"context"
	v1 "my-user-http-service/api/user/v1"
	"my-user-http-service/internal/dao"
	"my-user-http-service/internal/model/do"
)

func SignUp(ctx context.Context, req *v1.SignUpReq) (err error) {
	_, err = dao.User.Ctx(ctx).Data(do.User{
		Passport: req.Passport,
		Password: req.Password,
		Nickname: req.Nickname,
	}).Insert()
	return
}
