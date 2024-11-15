package postgres

import (
	"context"
	"database/sql"
	"server/internal/entities"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetById(ctx context.Context, id int64) (*entities.User, error) {
	user := entities.User{}
	query := `SELECT FROM users (id, name, surname, third_name, email, phone_number, birth_date, role) 
	WHERE id=$1`
	err := r.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Name, &user.Surname, &user.Email,
		&user.PhoneNumber, &user.BirthDate, &user.Role)
	if err != nil {
		return &entities.User{}, nil
	}

	return &user, nil
}

func (r *UserRepository) GetByData(ctx context.Context, email string, phoneNumber string) (*entities.User, error) {
	user := entities.User{}
	query := `SELECT id, name, surname, third_name, email, phone_number, birth_date, role, password FROM users  
WHERE email=$1 OR phone_number=$2`
	err := r.db.QueryRowContext(ctx, query, email, phoneNumber).Scan(&user.ID, &user.Name, &user.Surname, &user.ThirdName,
		&user.Email, &user.PhoneNumber, &user.BirthDate, &user.Role, &user.Password)
	if err != nil {
		return &entities.User{}, err
	}

	return &user, nil

}

func (r *UserRepository) Exists(ctx context.Context, phoneNumber string, email string) (bool, error) {
	var count int
	query := "SELECT COUNT(1) FROM users WHERE phone_number=$1 AND email=$2"

	err := r.db.QueryRowContext(ctx, query, phoneNumber, email).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *UserRepository) CreateUser(ctx context.Context, user *entities.User) (*entities.User, error) {
	var lastInsertId int
	query := `INSERT INTO users (name, surname, third_name, email, phone_number, password, birth_date, role)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`

	err := r.db.QueryRowContext(ctx, query, user.Name, user.Surname, user.ThirdName, user.Email,
		user.PhoneNumber, user.Password, user.BirthDate, user.Role).Scan(&lastInsertId)
	if err != nil {
		return &entities.User{}, err
	}
	user.ID = lastInsertId
	return user, nil
}
