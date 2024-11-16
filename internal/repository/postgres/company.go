package postgres

import (
	"context"
	"database/sql"
	"server/internal/entities"
)

type CompanyRepository struct {
	db *sql.DB
}

func NewCompanyRepository(db *sql.DB) *CompanyRepository {
	return &CompanyRepository{
		db: db,
	}
}

func (r *CompanyRepository) Create(ctx context.Context, company *entities.Company) (*entities.Company, error) {
	query := `
		INSERT INTO companies (name) VALUES ($1) RETURNING id;
	`

	err := r.db.QueryRowContext(ctx, query, company.Name).Scan(&company.ID)
	if err != nil {
		return nil, err
	}
	return company, nil
}
