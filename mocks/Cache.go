// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	utils "github.com/aasumitro/posbe/pkg/utils"
	mock "github.com/stretchr/testify/mock"
)

// Cache is an autogenerated mock type for the Cache type
type Cache struct {
	mock.Mock
}

// CacheFirstData provides a mock function with given fields: i
func (_m *Cache) CacheFirstData(i *utils.CacheDataSupplied) (interface{}, error) {
	ret := _m.Called(i)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(*utils.CacheDataSupplied) interface{}); ok {
		r0 = rf(i)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*utils.CacheDataSupplied) error); ok {
		r1 = rf(i)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewCache interface {
	mock.TestingT
	Cleanup(func())
}

// NewCache creates a new instance of Cache. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCache(t mockConstructorTestingTNewCache) *Cache {
	mock := &Cache{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
