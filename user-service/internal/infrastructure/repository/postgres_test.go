package repository_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/hossein-225/Library-Management/user-service/internal/domain"
	"github.com/hossein-225/Library-Management/user-service/internal/infrastructure/repository"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestPostgresUserRepository_RegisterUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	userRepo := repository.NewPostgresUserRepository(db)
	user := &domain.User{
		Name:     "Hossein",
		Email:    "hossein@example.com",
		Password: "password",
		Role:     "user",
	}

	mock.ExpectExec("INSERT INTO users").
		WithArgs(user.Name, user.Email, sqlmock.AnyArg(), user.Role).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = userRepo.RegisterUser(user)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestPostgresUserRepository_AuthenticateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	userRepo := repository.NewPostgresUserRepository(db)
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("secret"), bcrypt.DefaultCost)
	user := &domain.User{
		Name:     "Hossein",
		Email:    "hossein@example.com",
		Password: string(hashedPassword),
		Role:     "user",
	}

	mock.ExpectQuery("SELECT name, email, password, role FROM users").
		WithArgs(user.Email).
		WillReturnRows(sqlmock.NewRows([]string{"name", "email", "password", "role"}).
			AddRow(user.Name, user.Email, user.Password, user.Role))

	authenticatedUser, err := userRepo.AuthenticateUser(user.Email, "secret")
	assert.NoError(t, err)
	assert.Equal(t, user.Name, authenticatedUser.Name)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestPostgresUserRepository_GetUserProfile(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	userRepo := repository.NewPostgresUserRepository(db)
	user := &domain.User{
		Name:  "Hossein",
		Email: "hossein@example.com",
		Role:  "user",
	}

	mock.ExpectQuery("SELECT name, email, role FROM users").
		WithArgs(user.Email).
		WillReturnRows(sqlmock.NewRows([]string{"name", "email", "role"}).
			AddRow(user.Name, user.Email, user.Role))

	fetchedUser, err := userRepo.GetUserProfile(user.Email)
	assert.NoError(t, err)
	assert.Equal(t, user.Name, fetchedUser.Name)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestPostgresUserRepository_UpdateUserProfile(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	userRepo := repository.NewPostgresUserRepository(db)
	user := &domain.User{
		Name:  "hossein Updated",
		Email: "hossein@example.com",
	}

	mock.ExpectExec("UPDATE users SET name").
		WithArgs(user.Name, user.Email).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = userRepo.UpdateUserProfile(user)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
