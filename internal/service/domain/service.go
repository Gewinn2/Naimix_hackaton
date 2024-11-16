package domain

import (
	"server/internal/config"
	"server/internal/repository"
	"server/internal/service"
	"server/internal/service/domain/company"
	"server/internal/service/domain/cosmogram"
	"server/internal/service/domain/taro"
	"server/internal/service/domain/user"
)

func NewService(repositories *repository.Repository, conf *config.Config) *service.Service {
	return &service.Service{
		UserService:      user.NewUserService(repositories.User, conf),
		CompanyService:   company.NewCompanyService(repositories.Company, repositories.CompanyMember),
		TaroService:      taro.NewTaroService(repositories.TaroCard),
		CosmogramService: cosmogram.NewCosmogramService(),
	}
}
