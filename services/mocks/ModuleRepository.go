// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	models "github.com/aofiee/barroth/models"
	mock "github.com/stretchr/testify/mock"
)

// ModuleRepository is an autogenerated mock type for the ModuleRepository type
type ModuleRepository struct {
	mock.Mock
}

// CreateModule provides a mock function with given fields: m
func (_m *ModuleRepository) CreateModule(m *models.Modules) error {
	ret := _m.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Modules) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllRoles provides a mock function with given fields:
func (_m *ModuleRepository) GetAllRoles() ([]models.RoleItems, error) {
	ret := _m.Called()

	var r0 []models.RoleItems
	if rf, ok := ret.Get(0).(func() []models.RoleItems); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.RoleItems)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetModule provides a mock function with given fields: m, id
func (_m *ModuleRepository) GetModule(m *models.Modules, id string) error {
	ret := _m.Called(m, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Modules, string) error); ok {
		r0 = rf(m, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetModuleBySlug provides a mock function with given fields: m, method, slug
func (_m *ModuleRepository) GetModuleBySlug(m *models.Modules, method string, slug string) error {
	ret := _m.Called(m, method, slug)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Modules, string, string) error); ok {
		r0 = rf(m, method, slug)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetPermission provides a mock function with given fields: moduleID, roleID, exec
func (_m *ModuleRepository) SetPermission(moduleID uint, roleID uint, exec int) error {
	ret := _m.Called(moduleID, roleID, exec)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, uint, int) error); ok {
		r0 = rf(moduleID, roleID, exec)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateModule provides a mock function with given fields: m, id
func (_m *ModuleRepository) UpdateModule(m *models.Modules, id string) error {
	ret := _m.Called(m, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Modules, string) error); ok {
		r0 = rf(m, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
