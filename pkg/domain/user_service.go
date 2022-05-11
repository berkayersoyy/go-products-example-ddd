package domain

import "context"

type UserService interface {
	GetAllUsers() []User
	GetUserByID(id uint) User
	AddUser(user User) User
	GetUserByUsername(username string) User
	DeleteUser(User)
}

type UserServiceDynamoDb interface {
	Update(ctx context.Context, user User) error
	Find(ctx context.Context, id uint) (User, error)
	Insert(ctx context.Context, user User) error
	Delete(ctx context.Context, id uint) error
}
