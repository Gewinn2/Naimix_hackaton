package service

import (
	"server/internal/config"
	"server/internal/repository"
)

// TODO Дополнить для других сервисов
type Service struct {
	conf *config.Config
}

func NewService(repositories *repository.Repository, conf *config.Config) *Service {
	return &Service{}
}
