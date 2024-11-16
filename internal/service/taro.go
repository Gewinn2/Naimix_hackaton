package service

import (
	"context"
	"server/internal/entities"
	"server/internal/repository"
	"time"
)

type TaroService struct {
	repository repository.TaroCardRepository
	timeout    time.Duration
}

func NewTaroService(repository repository.TaroCardRepository) *TaroService {
	return &TaroService{
		repository: repository,
		timeout:    time.Duration(10) * time.Second,
	}
}

func (s *TaroService) GetById(c context.Context, id int) (*entities.TaroCard, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	card, err := s.repository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return card, nil
}
