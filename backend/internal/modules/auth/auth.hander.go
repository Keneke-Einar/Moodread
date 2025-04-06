package auth

import (
	"backend/internal/modules"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthenticationHandler struct {
	modules.BaseHandler
}

func NewAuthorizationHandler(service modules.Service, logger *zap.SugaredLogger) *AuthenticationHandler {
	authHandler := &AuthenticationHandler{}
	authHandler.SetService(service)
	authHandler.SetLogger(logger)
	return authHandler
}

func (h *AuthenticationHandler) Register(c *gin.Context) {
	c.JSON(200, gin.H{"message": "your account has been created"})
}

func (h *AuthenticationHandler) Login(c *gin.Context) {
	c.JSON(200, gin.H{"message": "you are logged in"})
}

func (h *AuthenticationHandler) Logout(c *gin.Context) {
	c.JSON(200, gin.H{"message": "you are logged out"})
}
