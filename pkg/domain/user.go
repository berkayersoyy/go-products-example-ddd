package domain

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// swagger:model User
type User struct {
	gorm.Model
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// swagger:model UserDTO
type UserDTO struct {
	ID       uint   `json:"id,string,omitempty"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func ToUser(userDTO UserDTO) User {

	return User{Username: userDTO.Username, Password: userDTO.Password}
}

func ToUserDTO(user User) UserDTO {
	return UserDTO{ID: user.ID, Username: user.Username, Password: user.Password}
}

func ToUserDTOs(users []User) []UserDTO {
	userdtos := make([]UserDTO, len(users))

	for i, itm := range users {
		userdtos[i] = ToUserDTO(itm)
	}

	return userdtos
}

type UserHandler interface {
	GetAllUsers(c *gin.Context)
	GetUserByID(c *gin.Context)
	AddUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type UserHandlerDynamoDb interface {
	Update(c *gin.Context)
	Find(c *gin.Context)
	Insert(c *gin.Context)
	Delete(c *gin.Context)
}
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
