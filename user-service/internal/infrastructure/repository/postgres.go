package repository

import (
	"database/sql"

	"github.com/hossein-225/Library-Management/user-service/internal/domain"
)

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) RegisterUser(user *domain.User) error {
	_, err := r.db.Exec("INSERT INTO users (id, name, email, password) VALUES ($1, $2, $3, $4)", user.ID, user.Name, user.Email, user.Password)
	return err
}

func (r *PostgresUserRepository) AuthenticateUser(email, password string) (*domain.User, error) {
	user := &domain.User{}
	err := r.db.QueryRow("SELECT id, name, email, password FROM users WHERE email=$1 AND password=$2", email, password).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *PostgresUserRepository) GetUserProfile(id string) (*domain.User, error) {
	user := &domain.User{}
	err := r.db.QueryRow("SELECT id, name, email FROM users WHERE id=$1", id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *PostgresUserRepository) UpdateUserProfile(user *domain.User) error {
	_, err := r.db.Exec("UPDATE users SET name=$1, email=$2 WHERE id=$3", user.Name, user.Email, user.ID)
	return err
}
