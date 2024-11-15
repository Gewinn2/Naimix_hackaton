package repository

import (
	"context"
	"server/internal/entities"
)

type UserRepository interface {
	GetById(context.Context, int64) (*entities.User, error)
	GetByData(context.Context, string, string) (*entities.User, error)
	Exists(context.Context, string, string) (bool, error)
	CreateUser(context.Context, *entities.User) (*entities.User, error)
}

type Repository struct {
	User UserRepository
}
