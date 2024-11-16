package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"server/internal/entities"
)

type TaroCardRepository struct {
	db *sql.DB
}

func NewTaroCardRepository(db *sql.DB) *TaroCardRepository {
	return &TaroCardRepository{
		db: db,
	}
}

func (r *TaroCardRepository) GetById(ctx context.Context, id int) (*entities.TaroCard, error) {

	log.Println("Entered DB to get Taro card")
	card := &entities.TaroCard{}
	query := `
		SELECT id, direct_meaning, reverse_meaning
		FROM taro_cards WHERE id=$1;

	`

	err := r.db.QueryRowContext(ctx, query, id).Scan(&card.ID, &card.DirectMeaning, &card.ReverseMeaning)
	if err != nil {
		if err == sql.ErrNoRows {

			return nil, fmt.Errorf("no card found with ID %d", id)
		}
		log.Printf("Error executing query: %v", err)
		return nil, err
	}
	return card, nil

}
