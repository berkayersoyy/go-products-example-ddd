package domain

import "context"

type UserRepository interface {
	GetAllUsers() []User
	GetUserByID(id uint) User
	GetUserByUsername(username string) User
	AddUser(user User) User
	DeleteUser(user User)
}
type UserRepositoryCtx interface {
	GetAllUsers(ctx context.Context) ([]User, error)
	GetUserByID(ctx context.Context, id uint) (User, error)
	GetUserByUsername(ctx context.Context, username string) (User, error)
	AddUser(ctx context.Context, user User) error
	DeleteUser(ctx context.Context, user User) error
}
