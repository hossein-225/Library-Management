package repository

import (
	"database/sql"

	"github.com/hossein-225/Library-Management/user-service/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) RegisterUser(user *domain.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = r.db.Exec("INSERT INTO users (name, email, password, role) VALUES ($1, $2, $3, $4, $5)", user.Name, user.Email, string(hashedPassword), user.Role)
	return err
}

func (r *PostgresUserRepository) AuthenticateUser(email, password string) (*domain.User, error) {
	user := &domain.User{}
	err := r.db.QueryRow("SELECT name, email, password, role FROM users WHERE email=$1", email).
		Scan(
			&user.Name,
			&user.Email,
			&user.Password,
			&user.Role)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *PostgresUserRepository) GetUserProfile(email string) (*domain.User, error) {
	user := &domain.User{}
	err := r.db.QueryRow("SELECT name, email, role FROM users WHERE email=$1", email).
		Scan(
			&user.Name,
			&user.Email,
			&user.Role)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *PostgresUserRepository) UpdateUserProfile(user *domain.User) error {
	_, err := r.db.Exec("UPDATE users SET name=$1 WHERE email=$2", user.Name, user.Email)
	return err
}
