// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	domain "github.com/berkayersoyy/go-products-example-ddd/pkg/domain"
	mock "github.com/stretchr/testify/mock"
)

// UserService is an autogenerated mock type for the UserService type
type UserService struct {
	mock.Mock
}

// AddUser provides a mock function with given fields: user
func (_m *UserService) AddUser(user domain.User) domain.User {
	ret := _m.Called(user)

	var r0 domain.User
	if rf, ok := ret.Get(0).(func(domain.User) domain.User); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	return r0
}

// DeleteUser provides a mock function with given fields: _a0
func (_m *UserService) DeleteUser(_a0 domain.User) {
	_m.Called(_a0)
}

// GetAllUsers provides a mock function with given fields:
func (_m *UserService) GetAllUsers() []domain.User {
	ret := _m.Called()

	var r0 []domain.User
	if rf, ok := ret.Get(0).(func() []domain.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.User)
		}
	}

	return r0
}

// GetUserByID provides a mock function with given fields: id
func (_m *UserService) GetUserByID(id uint) domain.User {
	ret := _m.Called(id)

	var r0 domain.User
	if rf, ok := ret.Get(0).(func(uint) domain.User); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	return r0
}

// GetUserByUsername provides a mock function with given fields: username
func (_m *UserService) GetUserByUsername(username string) domain.User {
	ret := _m.Called(username)

	var r0 domain.User
	if rf, ok := ret.Get(0).(func(string) domain.User); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	return r0
}
