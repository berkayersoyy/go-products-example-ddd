package domain

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//User Entity_user
// swagger:model User
type User struct {
	gorm.Model
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

//UserDTO Dto_user
// swagger:model UserDTO
type UserDTO struct {
	ID       uint   `json:"id,string,omitempty"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

//ToUser To_user
func ToUser(userDTO UserDTO) User {

	return User{Username: userDTO.Username, Password: userDTO.Password}
}

//ToUserDTO To_user_dto
func ToUserDTO(user User) UserDTO {
	return UserDTO{ID: user.ID, Username: user.Username, Password: user.Password}
}

//ToUserDTOs To_user_dtos
func ToUserDTOs(users []User) []UserDTO {
	userdtos := make([]UserDTO, len(users))

	for i, itm := range users {
		userdtos[i] = ToUserDTO(itm)
	}

	return userdtos
}

//UserHandler User_handler
type UserHandler interface {
	GetAllUsers(c *gin.Context)
	GetUserByID(c *gin.Context)
	AddUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

//UserHandlerDynamoDb User_handler_dynamodb
type UserHandlerDynamoDb interface {
	Update(c *gin.Context)
	Find(c *gin.Context)
	Insert(c *gin.Context)
	Delete(c *gin.Context)
}

//UserRepository User_repository
type UserRepository interface {
	GetAllUsers() []User
	GetUserByID(id uint) User
	GetUserByUsername(username string) User
	AddUser(user User) User
	DeleteUser(user User)
}

//UserRepositoryCtx User_repository_ctx
type UserRepositoryCtx interface {
	Update(ctx context.Context, user User) error
	Find(ctx context.Context, id uint) (User, error)
	Insert(ctx context.Context, user User) error
	Delete(ctx context.Context, id uint) error
}

//UserService User_service
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
