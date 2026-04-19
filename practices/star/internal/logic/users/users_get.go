package users

import (
	"context"
	"star/internal/dao"
	"star/internal/model/entity"
)

func (u *Users) GetUser(ctx context.Context, id int64) (user *entity.Users, err error) {
	user = &entity.Users{}
	err = dao.Users.Ctx(ctx).WherePri(id).Scan(user)
	return
}
