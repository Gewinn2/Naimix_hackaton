package repository

import (
	"context"
	"server/internal/entities"
)

type UserRepository interface {
	GetById(context.Context, int64) (*entities.User, error)
	GetByData(context.Context, string, string) (*entities.User, error)
	GetAll(context.Context) (*[]entities.User, error)
	Exists(context.Context, string, string) (bool, error)
	CreateUser(context.Context, *entities.User) (*entities.User, error)
}

type CompanyRepository interface {
	Create(context.Context, *entities.Company) (*entities.Company, error)
}

type CompanyMemberRepository interface {
	Create(context.Context, *entities.CompanyMember) (*entities.CompanyMember, error)
	GetAll(context.Context, int) ([]entities.CompanyMemberInfo, error)
}

type TaroCardRepository interface {
	GetCardsByCategoryID(context.Context, int) ([]entities.TaroCard, error)
	GetCategoryCount(context.Context, int) (int, error)
}

type Repository struct {
	User          UserRepository
	Company       CompanyRepository
	CompanyMember CompanyMemberRepository
	TaroCard      TaroCardRepository
}
