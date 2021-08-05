// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	models "github.com/aofiee/barroth/models"
	mock "github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: m
func (_m *UserRepository) CreateUser(m *models.Users) error {
	ret := _m.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Users) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateUserRole provides a mock function with given fields: m
func (_m *UserRepository) CreateUserRole(m *models.UserRoles) error {
	ret := _m.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.UserRoles) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteUsers provides a mock function with given fields: focus, id
func (_m *UserRepository) DeleteUsers(focus string, id []string) (int64, error) {
	ret := _m.Called(focus, id)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string, []string) int64); ok {
		r0 = rf(focus, id)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, []string) error); ok {
		r1 = rf(focus, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllUsers provides a mock function with given fields: m, keyword, sorting, sortField, page, limit, focus
func (_m *UserRepository) GetAllUsers(m *[]models.Users, keyword string, sorting string, sortField string, page string, limit string, focus string) (int64, error) {
	ret := _m.Called(m, keyword, sorting, sortField, page, limit, focus)

	var r0 int64
	if rf, ok := ret.Get(0).(func(*[]models.Users, string, string, string, string, string, string) int64); ok {
		r0 = rf(m, keyword, sorting, sortField, page, limit, focus)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*[]models.Users, string, string, string, string, string, string) error); ok {
		r1 = rf(m, keyword, sorting, sortField, page, limit, focus)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUser provides a mock function with given fields: m, uuid
func (_m *UserRepository) GetUser(m *models.Users, uuid string) error {
	ret := _m.Called(m, uuid)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Users, string) error); ok {
		r0 = rf(m, uuid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetUserByEmail provides a mock function with given fields: m, email
func (_m *UserRepository) GetUserByEmail(m *models.Users, email string) error {
	ret := _m.Called(m, email)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Users, string) error); ok {
		r0 = rf(m, email)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetUserRole provides a mock function with given fields: uid
func (_m *UserRepository) GetUserRole(uid uint) error {
	ret := _m.Called(uid)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(uid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HashPassword provides a mock function with given fields: user
func (_m *UserRepository) HashPassword(user *models.Users) error {
	ret := _m.Called(user)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Users) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RestoreUsers provides a mock function with given fields: id
func (_m *UserRepository) RestoreUsers(id []string) (int64, error) {
	ret := _m.Called(id)

	var r0 int64
	if rf, ok := ret.Get(0).(func([]string) int64); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: m, id
func (_m *UserRepository) UpdateUser(m *models.Users, id string) error {
	ret := _m.Called(m, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Users, string) error); ok {
		r0 = rf(m, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateUserRole provides a mock function with given fields: m, uid
func (_m *UserRepository) UpdateUserRole(m *models.UserRoles, uid uint) error {
	ret := _m.Called(m, uid)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.UserRoles, uint) error); ok {
		r0 = rf(m, uid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
