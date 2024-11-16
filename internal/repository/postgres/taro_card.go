package postgres

import (
	"context"
	"database/sql"
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

func (r *TaroCardRepository) GetCardsByCategoryID(ctx context.Context, id int) ([]entities.TaroCard, error) {

	var cards []entities.TaroCard
	query := `
		SELECT id, direct_meaning, reverse_meaning
		FROM taro_cards WHERE subcategory=$1;
	`
	rows, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		return cards, err
	}
	defer rows.Close()

	for rows.Next() {
		var card entities.TaroCard
		if err := rows.Scan(&card.ID, &card.DirectMeaning, &card.ReverseMeaning); err != nil {
			return cards, err
		}
		cards = append(cards, card)
	}
	if err := rows.Err(); err != nil {
		return cards, err
	}
	return cards, nil
}

func (r *TaroCardRepository) GetCategoryCount(ctx context.Context, categoryID int) (int, error) {
	query := `
		SELECT COUNT(*) AS card_count
		FROM taro_cards
		WHERE subcategory = $1;
	`

	var count int
	err := r.db.QueryRowContext(ctx, query, categoryID).Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}
	return count, nil
}
