package postgres

import (
	"context"
	"database/sql"
	"server/internal/entities"
)

type CompanyMemberRepository struct {
	db *sql.DB
}

func NewCompanyMemberRepository(db *sql.DB) *CompanyMemberRepository {
	return &CompanyMemberRepository{
		db: db,
	}
}

func (r *CompanyMemberRepository) Create(ctx context.Context, member *entities.CompanyMember) (*entities.CompanyMember, error) {
	query := `
		INSERT INTO company_members (user_id, company_id, position)
		VALUES ($1, $2, $3)
	`

	_, err := r.db.ExecContext(ctx, query, member.UserID, member.CompanyID, member.Position)
	if err != nil {
		return nil, err
	}
	return member, nil
}

func (r *CompanyMemberRepository) GetAll(ctx context.Context, companyID int) ([]entities.CompanyMemberInfo, error) {
	query := `
		SELECT 
			users.id AS user_id,
			users.name,
			users.surname,
			users.third_name,
			company_members.position
		FROM 
			company_members
		JOIN 
			users ON company_members.user_id = users.id
		WHERE 
			company_members.company_id = $1
	`

	rows, err := r.db.QueryContext(ctx, query, companyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var membersInfo []entities.CompanyMemberInfo
	for rows.Next() {
		var memberInfo entities.CompanyMemberInfo
		if err := rows.Scan(&memberInfo.UserID, &memberInfo.Name, &memberInfo.Surname, &memberInfo.ThirdName, &memberInfo.Position); err != nil {
			return nil, err
		}
		membersInfo = append(membersInfo, memberInfo)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return membersInfo, nil
}
