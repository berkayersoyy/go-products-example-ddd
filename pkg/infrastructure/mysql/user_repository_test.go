package mysql

//OBSOLETE
//import (
//	"github.com/berkayersoyy/go-products-example-ddd/pkg/domain"
//	"github.com/berkayersoyy/go-products-example-ddd/pkg/mocks"
//	"github.com/stretchr/testify/assert"
//	"github.com/stretchr/testify/mock"
//	"testing"
//	"time"
//)
//
//func TestUserRepository_GetAllUsersShouldReturnNotEmptyUserArray(t *testing.T) {
//	users := []domain.User{{Username: "test-username", Password: "test-pass", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: nil}}
//	mockRepo := mocks.UserRepository{}
//	mockRepo.On("GetAllUsers").Return(users)
//
//	resp := mockRepo.GetAllUsers()
//
//	assert.Equal(t, users, resp)
//	assert.NotEmpty(t, users)
//	assert.NotNil(t, resp)
//	mockRepo.AssertNumberOfCalls(t, "GetAllUsers", 1)
//}
//func TestUserRepository_GetAllUsersShouldReturnEmptyUserArray(t *testing.T) {
//	users := []domain.User{}
//	mockRepo := mocks.UserRepository{}
//	mockRepo.On("GetAllUsers").Return(users)
//
//	resp := mockRepo.GetAllUsers()
//
//	assert.Equal(t, users, resp)
//	assert.Empty(t, users)
//	assert.NotNil(t, resp)
//	mockRepo.AssertNumberOfCalls(t, "GetAllUsers", 1)
//}
//func TestUserRepository_GetUserByIDShouldReturnValidUser(t *testing.T) {
//	user := domain.User{Username: "test-username", Password: "test-pass", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: nil}
//	mockRepo := mocks.UserRepository{}
//	mockRepo.On("GetUserByID", mock.Anything).Return(user)
//
//	resp := mockRepo.GetUserByID(uint(1))
//
//	assert.Equal(t, user, resp)
//	assert.NotEmpty(t, resp)
//	assert.NotNil(t, resp)
//	mockRepo.AssertNumberOfCalls(t, "GetUserByID", 1)
//}
//func TestUserRepository_GetUserByIDShouldReturnEmptyUser(t *testing.T) {
//	user := domain.User{}
//	mockRepo := mocks.UserRepository{}
//	mockRepo.On("GetUserByID", mock.Anything).Return(user)
//
//	resp := mockRepo.GetUserByID(uint(0))
//
//	assert.Equal(t, user, resp)
//	assert.Empty(t, resp)
//	assert.NotNil(t, resp)
//	mockRepo.AssertNumberOfCalls(t, "GetUserByID", 1)
//}
//func TestUserRepository_GetUserByUsernameShouldReturnValidUser(t *testing.T) {
//	user := domain.User{Username: "test-username", Password: "test-pass", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: nil}
//	mockRepo := mocks.UserRepository{}
//	mockRepo.On("GetUserByUsername", mock.Anything).Return(user)
//
//	resp := mockRepo.GetUserByUsername("test-username")
//
//	assert.Equal(t, user, resp)
//	assert.NotEmpty(t, resp)
//	assert.NotNil(t, resp)
//	mockRepo.AssertNumberOfCalls(t, "GetUserByUsername", 1)
//}
//func TestUserRepository_GetUserByUsernameShouldReturnEmptyUser(t *testing.T) {
//	user := domain.User{}
//	mockRepo := mocks.UserRepository{}
//	mockRepo.On("GetUserByUsername", mock.Anything).Return(user)
//
//	resp := mockRepo.GetUserByUsername("test")
//
//	assert.Equal(t, user, resp)
//	assert.Empty(t, resp)
//	assert.NotNil(t, resp)
//	mockRepo.AssertNumberOfCalls(t, "GetUserByUsername", 1)
//}
//func TestUserRepository_AddUserShouldReturnValidUser(t *testing.T) {
//	user := domain.User{Username: "test-username", Password: "test-pass", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: nil}
//	mockRepo := mocks.UserRepository{}
//	mockRepo.On("AddUser", mock.Anything).Return(user)
//
//	resp := mockRepo.AddUser(user)
//
//	assert.Equal(t, user, resp)
//	assert.NotEmpty(t, resp)
//	assert.NotNil(t, resp)
//	mockRepo.AssertNumberOfCalls(t, "AddUser", 1)
//}
//func TestUserRepository_AddUserShouldReturnEmptyUser(t *testing.T) {
//	user := domain.User{}
//	mockRepo := mocks.UserRepository{}
//	mockRepo.On("AddUser", mock.Anything).Return(user)
//
//	resp := mockRepo.AddUser(user)
//
//	assert.Equal(t, user, resp)
//	assert.Empty(t, resp)
//	assert.NotNil(t, resp)
//	mockRepo.AssertNumberOfCalls(t, "AddUser", 1)
//}
//func TestUserRepository_DeleteUser(t *testing.T) {
//	user := domain.User{Username: "test-username", Password: "test-pass", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: nil}
//	mockRepo := mocks.UserRepository{}
//	mockRepo.On("DeleteUser", mock.Anything)
//
//	mockRepo.DeleteUser(user)
//
//	mockRepo.AssertNumberOfCalls(t, "DeleteUser", 1)
//}
