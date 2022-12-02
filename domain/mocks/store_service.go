// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/aasumitro/posbe/domain"
	mock "github.com/stretchr/testify/mock"

	utils "github.com/aasumitro/posbe/pkg/utils"
)

// IStoreService is an autogenerated mock type for the IStoreService type
type IStoreService struct {
	mock.Mock
}

// AddFloor provides a mock function with given fields: data
func (_m *IStoreService) AddFloor(data *domain.Floor) (*domain.Floor, *utils.ServiceError) {
	ret := _m.Called(data)

	var r0 *domain.Floor
	if rf, ok := ret.Get(0).(func(*domain.Floor) *domain.Floor); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Floor)
		}
	}

	var r1 *utils.ServiceError
	if rf, ok := ret.Get(1).(func(*domain.Floor) *utils.ServiceError); ok {
		r1 = rf(data)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*utils.ServiceError)
		}
	}

	return r0, r1
}

// AddTable provides a mock function with given fields: data
func (_m *IStoreService) AddTable(data *domain.Table) (*domain.Table, *utils.ServiceError) {
	ret := _m.Called(data)

	var r0 *domain.Table
	if rf, ok := ret.Get(0).(func(*domain.Table) *domain.Table); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Table)
		}
	}

	var r1 *utils.ServiceError
	if rf, ok := ret.Get(1).(func(*domain.Table) *utils.ServiceError); ok {
		r1 = rf(data)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*utils.ServiceError)
		}
	}

	return r0, r1
}

// DeleteFloor provides a mock function with given fields: data
func (_m *IStoreService) DeleteFloor(data *domain.Floor) *utils.ServiceError {
	ret := _m.Called(data)

	var r0 *utils.ServiceError
	if rf, ok := ret.Get(0).(func(*domain.Floor) *utils.ServiceError); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*utils.ServiceError)
		}
	}

	return r0
}

// DeleteTable provides a mock function with given fields: data
func (_m *IStoreService) DeleteTable(data *domain.Table) *utils.ServiceError {
	ret := _m.Called(data)

	var r0 *utils.ServiceError
	if rf, ok := ret.Get(0).(func(*domain.Table) *utils.ServiceError); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*utils.ServiceError)
		}
	}

	return r0
}

// EditFloor provides a mock function with given fields: data
func (_m *IStoreService) EditFloor(data *domain.Floor) (*domain.Floor, *utils.ServiceError) {
	ret := _m.Called(data)

	var r0 *domain.Floor
	if rf, ok := ret.Get(0).(func(*domain.Floor) *domain.Floor); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Floor)
		}
	}

	var r1 *utils.ServiceError
	if rf, ok := ret.Get(1).(func(*domain.Floor) *utils.ServiceError); ok {
		r1 = rf(data)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*utils.ServiceError)
		}
	}

	return r0, r1
}

// EditTable provides a mock function with given fields: data
func (_m *IStoreService) EditTable(data *domain.Table) (*domain.Table, *utils.ServiceError) {
	ret := _m.Called(data)

	var r0 *domain.Table
	if rf, ok := ret.Get(0).(func(*domain.Table) *domain.Table); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Table)
		}
	}

	var r1 *utils.ServiceError
	if rf, ok := ret.Get(1).(func(*domain.Table) *utils.ServiceError); ok {
		r1 = rf(data)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*utils.ServiceError)
		}
	}

	return r0, r1
}

// FloorList provides a mock function with given fields:
func (_m *IStoreService) FloorList() ([]*domain.Floor, *utils.ServiceError) {
	ret := _m.Called()

	var r0 []*domain.Floor
	if rf, ok := ret.Get(0).(func() []*domain.Floor); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Floor)
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

// FloorsWithTables provides a mock function with given fields:
func (_m *IStoreService) FloorsWithTables() ([]*domain.Floor, *utils.ServiceError) {
	ret := _m.Called()

	var r0 []*domain.Floor
	if rf, ok := ret.Get(0).(func() []*domain.Floor); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Floor)
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

// TableList provides a mock function with given fields:
func (_m *IStoreService) TableList() ([]*domain.Table, *utils.ServiceError) {
	ret := _m.Called()

	var r0 []*domain.Table
	if rf, ok := ret.Get(0).(func() []*domain.Table); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Table)
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

type mockConstructorTestingTNewIStoreService interface {
	mock.TestingT
	Cleanup(func())
}

// NewIStoreService creates a new instance of IStoreService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIStoreService(t mockConstructorTestingTNewIStoreService) *IStoreService {
	mock := &IStoreService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
