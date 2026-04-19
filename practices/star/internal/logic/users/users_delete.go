package users

import (
	"context"
	"star/internal/dao"
)

func (u *Users) DeleteUser(ctx context.Context, id int64) (err error) {
	_, err = u.GetUser(ctx, id)
	if err != nil {
		return
	}
	_, err = dao.Users.Ctx(ctx).WherePri(id).Delete()
	return
}
