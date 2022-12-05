// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/aasumitro/posbe/domain"
	mock "github.com/stretchr/testify/mock"

	utils "github.com/aasumitro/posbe/pkg/utils"
)

// IAccountService is an autogenerated mock type for the IAccountService type
type IAccountService struct {
	mock.Mock
}

// AddRole provides a mock function with given fields: data
func (_m *IAccountService) AddRole(data *domain.Role) (*domain.Role, *utils.ServiceError) {
	ret := _m.Called(data)

	var r0 *domain.Role
	if rf, ok := ret.Get(0).(func(*domain.Role) *domain.Role); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Role)
		}
	}

	var r1 *utils.ServiceError
	if rf, ok := ret.Get(1).(func(*domain.Role) *utils.ServiceError); ok {
		r1 = rf(data)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*utils.ServiceError)
		}
	}

	return r0, r1
}

// AddUser provides a mock function with given fields: data
func (_m *IAccountService) AddUser(data *domain.User) (*domain.User, *utils.ServiceError) {
	ret := _m.Called(data)

	var r0 *domain.User
	if rf, ok := ret.Get(0).(func(*domain.User) *domain.User); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	var r1 *utils.ServiceError
	if rf, ok := ret.Get(1).(func(*domain.User) *utils.ServiceError); ok {
		r1 = rf(data)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*utils.ServiceError)
		}
	}

	return r0, r1
}

// DeleteRole provides a mock function with given fields: data
func (_m *IAccountService) DeleteRole(data *domain.Role) *utils.ServiceError {
	ret := _m.Called(data)

	var r0 *utils.ServiceError
	if rf, ok := ret.Get(0).(func(*domain.Role) *utils.ServiceError); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*utils.ServiceError)
		}
	}

	return r0
}

// DeleteUser provides a mock function with given fields: data
func (_m *IAccountService) DeleteUser(data *domain.User) *utils.ServiceError {
	ret := _m.Called(data)

	var r0 *utils.ServiceError
	if rf, ok := ret.Get(0).(func(*domain.User) *utils.ServiceError); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*utils.ServiceError)
		}
	}

	return r0
}

// EditRole provides a mock function with given fields: data
func (_m *IAccountService) EditRole(data *domain.Role) (*domain.Role, *utils.ServiceError) {
	ret := _m.Called(data)

	var r0 *domain.Role
	if rf, ok := ret.Get(0).(func(*domain.Role) *domain.Role); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Role)
		}
	}

	var r1 *utils.ServiceError
	if rf, ok := ret.Get(1).(func(*domain.Role) *utils.ServiceError); ok {
		r1 = rf(data)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*utils.ServiceError)
		}
	}

	return r0, r1
}

// EditUser provides a mock function with given fields: data
func (_m *IAccountService) EditUser(data *domain.User) (*domain.User, *utils.ServiceError) {
	ret := _m.Called(data)

	var r0 *domain.User
	if rf, ok := ret.Get(0).(func(*domain.User) *domain.User); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	var r1 *utils.ServiceError
	if rf, ok := ret.Get(1).(func(*domain.User) *utils.ServiceError); ok {
		r1 = rf(data)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*utils.ServiceError)
		}
	}

	return r0, r1
}

// RoleList provides a mock function with given fields:
func (_m *IAccountService) RoleList() ([]*domain.Role, *utils.ServiceError) {
	ret := _m.Called()

	var r0 []*domain.Role
	if rf, ok := ret.Get(0).(func() []*domain.Role); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Role)
		}
	}

	var r1 *utils.ServiceError
	if rf, ok := ret.Get(1).(func() *utils.ServiceError); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*utils.ServiceError)
		}
	}

	return r0, r1
}

// ShowUser provides a mock function with given fields: id
func (_m *IAccountService) ShowUser(id int) (*domain.User, *utils.ServiceError) {
	ret := _m.Called(id)

	var r0 *domain.User
	if rf, ok := ret.Get(0).(func(int) *domain.User); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	var r1 *utils.ServiceError
	if rf, ok := ret.Get(1).(func(int) *utils.ServiceError); ok {
		r1 = rf(id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*utils.ServiceError)
		}
	}

	return r0, r1
}

// UserList provides a mock function with given fields:
func (_m *IAccountService) UserList() ([]*domain.User, *utils.ServiceError) {
	ret := _m.Called()

	var r0 []*domain.User
	if rf, ok := ret.Get(0).(func() []*domain.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.User)
		}
	}

	var r1 *utils.ServiceError
	if rf, ok := ret.Get(1).(func() *utils.ServiceError); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*utils.ServiceError)
		}
	}

	return r0, r1
}

// VerifyUserCredentials provides a mock function with given fields: username, password
func (_m *IAccountService) VerifyUserCredentials(username string, password string) (interface{}, *utils.ServiceError) {
	ret := _m.Called(username, password)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string, string) interface{}); ok {
		r0 = rf(username, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 *utils.ServiceError
	if rf, ok := ret.Get(1).(func(string, string) *utils.ServiceError); ok {
		r1 = rf(username, password)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*utils.ServiceError)
		}
	}

	return r0, r1
}

type mockConstructorTestingTNewIAccountService interface {
	mock.TestingT
	Cleanup(func())
}

// NewIAccountService creates a new instance of IAccountService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIAccountService(t mockConstructorTestingTNewIAccountService) *IAccountService {
	mock := &IAccountService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
