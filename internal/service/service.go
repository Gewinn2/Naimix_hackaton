package service

import (
	"context"
	"server/internal/config"
	"server/internal/entities"
	"server/internal/repository"
)

type User interface {
	CreateUser(context.Context, *entities.CreateUserRequest) (*entities.CreateUserResponse, error)
	Login(context.Context, *entities.LoginUserRequest) (*entities.LoginUserResponse, error)
	GetAll(ctx context.Context) (*[]entities.User, error)
}

type Company interface {
	CreateCompany(context.Context, *entities.CreateCompanyRequest) (*entities.CreateCompanyResponse, error)
	GetAllMembers(context.Context, *entities.GetCompanyMembersInfoRequest) ([]entities.CompanyMemberInfo, error)
}

type Taro interface {
	GetById(ctx context.Context, id int) (*entities.TaroCard, error)
}

// TODO Дополнить для других сервисов
type Service struct {
	UserService    *UserService
	CompanyService *CompanyService
	TaroService    *TaroService
	conf           *config.Config
}

func NewService(repositories *repository.Repository, conf *config.Config) *Service {
	return &Service{
		UserService:    NewUserService(repositories.User, conf),
		CompanyService: NewCompanyService(repositories.Company, repositories.CompanyMember),
		TaroService:    NewTaroService(repositories.TaroCard),
	}
}
