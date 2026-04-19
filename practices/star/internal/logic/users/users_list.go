package users

import (
	"context"
	"star/internal/dao"
	"star/internal/model/entity"
)

func (u *Users) GetUserList(ctx context.Context) (list []*entity.Users, err error) {
	err = dao.Users.Ctx(ctx).Scan(&list)
	return
}
