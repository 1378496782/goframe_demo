package user

import (
	"context"
	"my-user-http-service/internal/model/entity"
)

func (s *Service) GetProfile(ctx context.Context) *entity.User {
	return s.sessionSvc.GetUser(ctx)
}
