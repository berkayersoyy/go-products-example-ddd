package domain

type UserRepository interface {
	GetAllUsers() []User
	GetUserByID(id uint) User
	GetUserByUsername(username string) User
	AddUser(user User) User
	DeleteUser(user User)
}
