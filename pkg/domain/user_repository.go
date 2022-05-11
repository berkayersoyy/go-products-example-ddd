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
	Update(ctx context.Context, user User) error
	Find(ctx context.Context, id uint) (User, error)
	Insert(ctx context.Context, user User) error
	Delete(ctx context.Context, id uint) error
}
