package service

import (
	"context"
	"server/internal/config"
	"server/internal/entities"
)

type UserService interface {
	CreateUser(context.Context, *entities.CreateUserRequest) (*entities.CreateUserResponse, error)
	Login(context.Context, *entities.LoginUserRequest) (*entities.LoginUserResponse, error)
	GetAll(ctx context.Context) (*[]entities.User, error)
}

type CompanyService interface {
	CreateCompany(context.Context, *entities.CreateCompanyRequest) (*entities.CreateCompanyResponse, error)
	GetAllMembers(context.Context, *entities.GetCompanyMembersInfoRequest) ([]entities.CompanyMemberInfo, error)
}

type TaroService interface {
	GetById(ctx context.Context, id int) ([]entities.TaroCard, error)
}

type CosmogramService interface {
	Get(entities.CosmogramRequestBody) (*entities.CosmogramResponseBody, error)
}

type Service struct {
	UserService      UserService
	CompanyService   CompanyService
	TaroService      TaroService
	CosmogramService CosmogramService
	conf             *config.Config
}
