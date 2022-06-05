package application

import "github.com/berkayersoyy/go-products-example-ddd/pkg/domain"

type userService struct {
	UserRepository domain.UserRepository
}

//ProvideUserService Provide user service
func ProvideUserService(u domain.UserRepository) domain.UserService {
	return &userService{UserRepository: u}
}
func (u *userService) GetAllUsers() []domain.User {
	return u.UserRepository.GetAllUsers()
}
func (u *userService) GetUserByID(id uint) domain.User {
	return u.UserRepository.GetUserByID(id)
}
func (u *userService) AddUser(user domain.User) domain.User {
	return u.UserRepository.AddUser(user)
}
func (u *userService) GetUserByUsername(username string) domain.User {
	return u.UserRepository.GetUserByUsername(username)
}
func (u *userService) DeleteUser(user domain.User) {
	u.UserRepository.DeleteUser(user)
}
