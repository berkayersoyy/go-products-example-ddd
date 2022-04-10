package domain

import "github.com/jinzhu/gorm"

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
	return UserDTO{ID: user.ID, Username: user.Password, Password: user.Password}
}

func ToUserDTOs(users []User) []UserDTO {
	userdtos := make([]UserDTO, len(users))

	for i, itm := range users {
		userdtos[i] = ToUserDTO(itm)
	}

	return userdtos
}
