package domain

type UserRepository interface {
	RegisterUser(user *User) error
	AuthenticateUser(email, password string) (*User, error)
	GetUserProfile(id string) (*User, error)
	UpdateUserProfile(user *User) error
}
