package mysql

//type userRepository struct {
//	DB *gorm.DB
//}

//OBSOLETE
//ProvideUserRepository Provide user repository
//func ProvideUserRepository(DB *gorm.DB) domain.UserRepository {
//	return &userRepository{DB: DB}
//}
//func (u *userRepository) GetAllUsers() []domain.User {
//	var users []domain.User
//	u.DB.Find(&users)
//
//	return users
//}
//
//func (u *userRepository) GetUserByID(id uint) domain.User {
//	var user domain.User
//	u.DB.First(&user, id)
//
//	return user
//}
//func (u *userRepository) GetUserByUsername(username string) domain.User {
//	var user domain.User
//	u.DB.Where("username = ?", username).First(&user)
//
//	return user
//}
//
//func (u *userRepository) AddUser(user domain.User) domain.User {
//	u.DB.Save(&user)
//
//	return user
//}
//
//func (u *userRepository) DeleteUser(user domain.User) {
//	u.DB.Delete(&user)
//}
