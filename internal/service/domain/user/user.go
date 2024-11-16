package user

import (
	"context"
	"errors"
	"fmt"
	"server/internal/config"
	"server/internal/entities"
	"server/internal/repository"
	"server/pkg"
	"server/util"
	"strconv"
	"time"
)

type UserService struct {
	repository repository.UserRepository
	timeout    time.Duration
	conf       *config.Config
}

func NewUserService(repository repository.UserRepository, conf *config.Config) *UserService {
	return &UserService{
		repository: repository,
		timeout:    time.Duration(10) * time.Second,
		conf:       conf,
	}
}

func (s *UserService) CreateUser(c context.Context, request *entities.CreateUserRequest) (*entities.CreateUserResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	exists, err := s.repository.Exists(ctx, request.PhoneNumber, request.Email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("user already exists")
	}

	hashedPassword, err := util.HashPassword(request.Password)
	if err != nil {
		return nil, err
	}

	u := &entities.User{
		Name:        request.Name,
		Surname:     request.Surname,
		ThirdName:   request.ThirdName,
		PhoneNumber: request.PhoneNumber,
		Email:       request.Email,
		Password:    hashedPassword,
		BirthDate:   request.BirthDate,
		Role:        request.Role,
	}

	r, err := s.repository.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}

	TokenExpiration, err := strconv.Atoi(s.conf.Application.TokenExpiration)
	if err != nil {
		return &entities.CreateUserResponse{}, errors.New("wrong data")
	}
	accessToken, err := pkg.GenerateAccessToken(u.ID, TokenExpiration,
		s.conf.Application.SigningKey)
	if err != nil {
		return &entities.CreateUserResponse{}, err
	}

	refreshToken, err := pkg.GenerateRefreshToken(u.ID, s.conf.Application.SigningKey)
	if err != nil {
		return &entities.CreateUserResponse{}, err
	}

	res := &entities.CreateUserResponse{
		ID:           r.ID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Email:        r.Email,
	}

	return res, nil

}

func (s *UserService) Login(c context.Context, request *entities.LoginUserRequest) (*entities.LoginUserResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	u, err := s.repository.GetByData(ctx, request.Email, request.PhoneNumber)

	if err != nil {
		return &entities.LoginUserResponse{}, errors.New("wrong data")
	}

	err = util.CheckPassword(request.Password, u.Password)
	if err != nil {
		return &entities.LoginUserResponse{}, errors.New("jopa")
	}

	TokenExpiration, err := strconv.Atoi(s.conf.Application.TokenExpiration)
	accessToken, err := pkg.GenerateAccessToken(u.ID, TokenExpiration, s.conf.Application.SigningKey)
	if err != nil {
		return &entities.LoginUserResponse{}, err
	}

	refreshToken, err := pkg.GenerateRefreshToken(u.ID, s.conf.Application.SigningKey)
	if err != nil {
		return &entities.LoginUserResponse{}, err
	}

	return &entities.LoginUserResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Email:        u.Email,
		ID:           u.ID,
	}, nil
}

func (s *UserService) GetAll(ctx context.Context) (*[]entities.User, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	users, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting all users: %w", err)
	}

	return users, nil
}