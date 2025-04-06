package modules

import (
	"context"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository interface {
	SetDB(db *gorm.DB)
	GetDB() *gorm.DB
}

type Service interface {
	SetRepo(repo Repository)
	GetRepo() Repository
}

type Handler interface {
	SetService(service Service)
	GetService() Service
}

type BaseRepo struct {
	db             *gorm.DB
	defaultContext context.Context
}

func (r *BaseRepo) SetDB(db *gorm.DB) {
	r.db = db
}

func (r *BaseRepo) GetDB() *gorm.DB {
	return r.db
}

func (r *BaseRepo) CloneWithTransaction(tx *gorm.DB) Repository {
	clone := *r
	clone.SetDB(tx)
	return &clone
}

type BaseService struct {
	repo Repository
	BaseLogger
}

func (s *BaseService) SetRepo(repo Repository) {
	s.repo = repo
}

func (s *BaseService) GetRepo() Repository {
	return s.repo
}

type BaseHandler struct {
	service Service
	BaseLogger
}

func (h *BaseHandler) SetService(service Service) {
	h.service = service
}

func (h *BaseHandler) GetService() Service {
	return h.service
}

type BaseLogger struct {
	logger *zap.SugaredLogger
}

func (s *BaseLogger) SetLogger(logger *zap.SugaredLogger) {
	s.logger = logger
}

func (s *BaseLogger) GetLogger() *zap.SugaredLogger {
	return s.logger
}
