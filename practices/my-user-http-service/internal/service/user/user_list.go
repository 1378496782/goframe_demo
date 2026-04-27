package user

import (
	"context"
	"my-user-http-service/internal/dao"
	"my-user-http-service/internal/model/entity"
)

func (s *Service) GetList(ctx context.Context) (users []*entity.User, err error) {
	err = dao.User.Ctx(ctx).Scan(&users)
	return
}
