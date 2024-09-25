package domain

type AuthRepository interface {
	ValidateToken(token string) (string, error)
}
