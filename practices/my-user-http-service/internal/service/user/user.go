package user

import (
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
