package postgres

import (
	"database/sql"
	"server/internal/repository"
)

func NewRepository(db *sql.DB) *repository.Repository {
	return &repository.Repository{
		User:          NewUserRepository(db),
		Company:       NewCompanyRepository(db),
		CompanyMember: NewCompanyMemberRepository(db),
		TaroCard:      NewTaroCardRepository(db),
	}
	//TODO дополнить
}
