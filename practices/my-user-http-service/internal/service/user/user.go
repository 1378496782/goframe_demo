package user

import (
	"context"

	"my-user-http-service/internal/model/entity"
	"my-user-http-service/internal/service/bizctx"
	"my-user-http-service/internal/service/session"
)

// Service provides user-related business logic.
type Service struct {
	bizCtxSvc  *bizctx.Service  // Business context service.
	sessionSvc *session.Service // Session service.
}

// New creates and returns a new Service instance.
func New() *Service {
	return &Service{
		bizCtxSvc:  bizctx.New(),
		sessionSvc: session.New(),
	}
}

// SignOut removes the session for current signed-in user.
func (s *Service) SignOut(ctx context.Context) error {
	return s.sessionSvc.RemoveUser(ctx)
}

// IsSignedIn checks and returns whether current user is already signed-in.
func (s *Service) IsSignedIn(ctx context.Context) bool {
	if v := s.bizCtxSvc.Get(ctx); v != nil && v.User != nil {
		return true
	}
	return false
}

// GetProfile retrieves and returns current user info in session.
func (s *Service) GetProfile(ctx context.Context) *entity.User {
	return s.sessionSvc.GetUser(ctx)
}
