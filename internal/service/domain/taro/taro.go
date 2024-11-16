package taro

import (
	"context"
	"log"
	"math/rand"
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

func (s *TaroService) GetCard(c context.Context, id int) (*entities.TaroCard, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()
	cards, err := s.repository.GetCardsByCategoryID(ctx, id)
	if err != nil {
		log.Printf("Error fetching card from repository: %v", err)
		return nil, err
	}
	randomSource := rand.NewSource(time.Now().UnixNano())
	random := rand.New(randomSource)
	randomIndex := random.Intn(len(cards))
	return &cards[randomIndex], nil
}
