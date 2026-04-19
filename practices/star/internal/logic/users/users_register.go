package users

import (
	"context"
	"star/internal/dao"
	"star/internal/model/do"

	"github.com/gogf/gf/v2/errors/gerror"
)

func (u *Users) Register(ctx context.Context, username, password, email string) (err error) {
	err = u.CheckUsername(ctx, username)
	if err != nil {
		return
	}
	_, err = dao.Users.Ctx(ctx).Data(do.Users{
		Username: username,
		Password: u.encryptPassword(password),
		Email:    email,
	}).Insert()
	return
}

func (u *Users) CheckUsername(ctx context.Context, username string) (err error) {
	count, err := dao.Users.Ctx(ctx).Where("username = ?", username).Count()
	if err != nil {
		return
	}
	if count > 0 {
		return gerror.New("用户已存在")
	}
	return nil
}
