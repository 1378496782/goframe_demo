package user

import (
	"context"
)

func (s *Service) SignOut(ctx context.Context) error {
	return s.sessionSvc.RemoveUser(ctx)
}
