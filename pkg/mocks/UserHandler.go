// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"
)

// UserHandler is an autogenerated mock type for the UserHandler type
type UserHandler struct {
	mock.Mock
}

// AddUser provides a mock function with given fields: c
func (_m *UserHandler) AddUser(c *gin.Context) {
	_m.Called(c)
}

// DeleteUser provides a mock function with given fields: c
func (_m *UserHandler) DeleteUser(c *gin.Context) {
	_m.Called(c)
}

// GetAllUsers provides a mock function with given fields: c
func (_m *UserHandler) GetAllUsers(c *gin.Context) {
	_m.Called(c)
}

// GetUserByID provides a mock function with given fields: c
func (_m *UserHandler) GetUserByID(c *gin.Context) {
	_m.Called(c)
}

// UpdateUser provides a mock function with given fields: c
func (_m *UserHandler) UpdateUser(c *gin.Context) {
	_m.Called(c)
}
