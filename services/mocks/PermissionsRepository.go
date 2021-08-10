// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	models "github.com/aofiee/barroth/models"
	mock "github.com/stretchr/testify/mock"
)

// PermissionsRepository is an autogenerated mock type for the PermissionsRepository type
type PermissionsRepository struct {
	mock.Mock
}

// SetPermissions provides a mock function with given fields: m
func (_m *PermissionsRepository) SetPermissions(m *models.Permissions) error {
	ret := _m.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Permissions) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
