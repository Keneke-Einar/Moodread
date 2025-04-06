package auth

import (
	"backend/internal/modules"
	"gorm.io/gorm"
)

type AuthenticationRepository interface {
	modules.Repository
}

type authenticationRepository struct {
	modules.BaseRepo
}

func NewAuthenticationRepository(db *gorm.DB) AuthenticationRepository {
	authRepo := &authenticationRepository{}
	authRepo.SetDB(db)
	return &authenticationRepository{}
}
