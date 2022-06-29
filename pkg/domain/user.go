package domain

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

//User Entity_user
// swagger:model User
type User struct {
	ID        string     `gorm:"primary_key" json:"Id"`
	UUID      string     `json:"UUID"`
	Username  string     `json:"Username" validate:"required"`
	Password  string     `json:"Password" validate:"required"`
	CreatedAt time.Time  `json:"CreatedAt"`
	UpdatedAt time.Time  `json:"UpdatedAt"`
	DeletedAt *time.Time `json:"DeletedAt"`
}

//UserDTO Dto_user
// swagger:model UserDTO
type UserDTO struct {
	UUID     string `json:"uuid"`
	ID       string `json:"id,string,omitempty"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

//ToUser To_user
func ToUser(userDTO UserDTO) User {

	return User{Username: userDTO.Username, Password: userDTO.Password}
}

//ToUserDTO To_user_dto
func ToUserDTO(user User) UserDTO {
	return UserDTO{UUID: user.UUID, ID: user.ID, Username: user.Username, Password: user.Password}
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
//type UserHandler interface {
//	GetAllUsers(c *gin.Context)
//	GetUserByID(c *gin.Context)
//	AddUser(c *gin.Context)
//	UpdateUser(c *gin.Context)
//	DeleteUser(c *gin.Context)
//}

//UserHandler User_handler_dynamodb
type UserHandler interface {
	Update(c *gin.Context)
	FindByUUID(c *gin.Context)
	Insert(c *gin.Context)
	Delete(c *gin.Context)
	FindByUsername(c *gin.Context)
}

//UserRepository User_repository
//type UserRepository interface {
//	GetAllUsers() []User
//	GetUserByID(id uint) User
//	GetUserByUsername(username string) User
//	AddUser(user User) User
//	DeleteUser(user User)
//}

//UserRepository User_repository_Dynamodb
type UserRepository interface {
	Update(ctx context.Context, user User) error
	FindByUUID(ctx context.Context, id string) (User, error)
	Insert(ctx context.Context, user User) error
	Delete(ctx context.Context, id string) error
	FindByUsername(ctx context.Context, username string) (User, error)
	CreateTable(ctx context.Context) error
}

//UserService User_service
//type UserService interface {
//	GetAllUsers() []User
//	GetUserByID(id uint) User
//	AddUser(user User) User
//	GetUserByUsername(username string) User
//	DeleteUser(User)
//}

//UserService User_service_dynamodb
type UserService interface {
	Update(ctx context.Context, user User) error
	FindByUUID(ctx context.Context, id string) (User, error)
	Insert(ctx context.Context, user User) error
	Delete(ctx context.Context, uuid string) error
	FindByUsername(ctx context.Context, username string) (User, error)
	CreateTable(ctx context.Context) error
}
