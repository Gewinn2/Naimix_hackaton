package taro

import (
	"context"
	"log"
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

func (s *TaroService) GetById(c context.Context, id int) ([]entities.TaroCard, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	log.Println("Calling repository to get Taro card")
	card, err := s.repository.GetCardsByCategoryID(ctx, id)
	if err != nil {
		log.Printf("Error fetching card from repository: %v", err)
		return nil, err
	}
	return card, nil

}
