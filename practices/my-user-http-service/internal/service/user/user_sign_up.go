package user

import (
	"context"
	"my-user-http-service/internal/dao"
	"my-user-http-service/internal/model/do"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
)

// CreateInput defines the input for Create function.
type CreateInput struct {
	Passport string
	Password string
	Nickname string
}

// Create creates user account.
func (s *Service) Create(ctx context.Context, in CreateInput) (err error) {
	// If Nickname is not specified, it then uses Passport as its default Nickname.
	if in.Nickname == "" {
		in.Nickname = in.Passport
	}

	return dao.User.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// Passport checks.
		count, err := dao.User.Ctx(ctx).Where(do.User{
			Passport: in.Passport,
		}).Count()
		if err != nil {
			return err
		}
		if count > 0 {
			return gerror.Newf(`Passport "%s" is already taken by others`, in.Passport)
		}

		// Nickname checks.
		count, err = dao.User.Ctx(ctx).Where(do.User{
			Nickname: in.Nickname,
		}).Count()
		if err != nil {
			return err
		}
		if count > 0 {
			return gerror.Newf(`Nickname "%s" is already taken by others`, in.Nickname)
		}

		_, err = dao.User.Ctx(ctx).Data(do.User{
			Passport: in.Passport,
			Password: in.Password,
			Nickname: in.Nickname,
		}).Insert()
		return err
	})
}
