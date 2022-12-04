// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/aasumitro/posbe/domain"
	mock "github.com/stretchr/testify/mock"

	utils "github.com/aasumitro/posbe/pkg/utils"
)

// IStorePrefService is an autogenerated mock type for the IStorePrefService type
type IStorePrefService struct {
	mock.Mock
}

// StoreSettings provides a mock function with given fields:
func (_m *IStorePrefService) AllPrefs() (*domain.StoreSetting, *utils.ServiceError) {
	ret := _m.Called()

	var r0 *domain.StoreSetting
	if rf, ok := ret.Get(0).(func() *domain.StoreSetting); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.StoreSetting)
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

// UpdateStorePrefs provides a mock function with given fields: key, value
func (_m *IStorePrefService) UpdatePrefs(key string, value string) (*domain.StoreSetting, *utils.ServiceError) {
	ret := _m.Called(key, value)

	var r0 *domain.StoreSetting
	if rf, ok := ret.Get(0).(func(string, string) *domain.StoreSetting); ok {
		r0 = rf(key, value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.StoreSetting)
		}
	}

	var r1 *utils.ServiceError
	if rf, ok := ret.Get(1).(func(string, string) *utils.ServiceError); ok {
		r1 = rf(key, value)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*utils.ServiceError)
		}
	}

	return r0, r1
}

type mockConstructorTestingTNewIStorePrefService interface {
	mock.TestingT
	Cleanup(func())
}

// NewIStorePrefService creates a new instance of IStorePrefService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIStorePrefService(t mockConstructorTestingTNewIStorePrefService) *IStorePrefService {
	mock := &IStorePrefService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}