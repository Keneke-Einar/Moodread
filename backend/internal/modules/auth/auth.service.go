package auth

import (
	"backend/internal/modules"
	"go.uber.org/zap"
)

type AuthorizationService interface {
	modules.Service
}

type authorizationService struct {
	modules.BaseService
}

func NewAuthorizationService(repo AuthenticationRepository, logger *zap.SugaredLogger) AuthorizationService {
	service := &authorizationService{}
	service.SetRepo(repo)
	service.SetLogger(logger)
	return service
}
