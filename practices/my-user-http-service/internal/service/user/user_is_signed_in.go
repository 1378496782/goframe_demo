package user

import "context"

func (s *Service) IsSignedIn(ctx context.Context) bool {
	if v := s.bizCtxSvc.Get(ctx); v != nil && v.User != nil {
		return true
	}
	return false
}
