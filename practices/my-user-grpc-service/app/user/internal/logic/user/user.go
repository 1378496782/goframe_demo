package user

import (
	"context"
	"my-user-grpc-service/app/user/api/pbentity"
	v1 "my-user-grpc-service/app/user/api/user/v1"
	"my-user-grpc-service/app/user/internal/dao"
)

func GetList(ctx context.Context) (users []*pbentity.User, err error) {
	err = dao.User.Ctx(ctx).Scan(&users)
	return
}

func GetById(ctx context.Context, uid uint64) (user *pbentity.User, err error) {
	err = dao.User.Ctx(ctx).WherePri(uid).Scan(&user)
	return
}

func Create(ctx context.Context, req *v1.CreateReq) (err error) {
	_, err = dao.User.Ctx(ctx).Data(pbentity.User{
		Passport: req.Passport,
		Password: req.Password,
		Nickname: req.Nickname,
	}).Insert()
	return
}

func Delete(ctx context.Context, uid uint64) (err error) {
	_, err = dao.User.Ctx(ctx).WherePri(uid).Delete()
	return
}
