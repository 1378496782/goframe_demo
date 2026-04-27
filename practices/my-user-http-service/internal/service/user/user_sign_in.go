package user

import (
	"context"
	"my-user-http-service/internal/dao"
	"my-user-http-service/internal/model/do"
	"my-user-http-service/internal/model/entity"
	"my-user-http-service/internal/service/bizctx"

	"github.com/gogf/gf/v2/errors/gerror"
)

type SignInInput struct {
	Passport string
	Password string
}

func (s *Service) SignIn(ctx context.Context, in SignInInput) (err error) {
	var (
		user = &entity.User{}
		orm  = dao.User.Ctx(ctx)
	)
	err = orm.Where(do.User{
		Passport: in.Passport,
		Password: in.Password,
	}).Scan(user)
	if err != nil {
		return err
	}
	if user.Id == 0 {
		return gerror.New("用户不存在")
	}
	err = s.sessionSvc.SetUser(ctx, user)
	if err != nil {
		return err
	}
	s.bizCtxSvc.SetUser(ctx, &bizctx.User{
		Id:       user.Id,
		Passport: user.Passport,
		Nickname: user.Nickname,
	})
	return nil
}
