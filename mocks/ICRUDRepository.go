// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/aasumitro/posbe/pkg/model"
	mock "github.com/stretchr/testify/mock"
)

// ICRUDRepository is an autogenerated mock type for the ICRUDRepository type
type ICRUDRepository[T interface{}] struct {
	mock.Mock
}

// All provides a mock function with given fields: ctx
func (_m *ICRUDRepository[T]) All(ctx context.Context) ([]*T, error) {
	ret := _m.Called(ctx)

	var r0 []*T
	if rf, ok := ret.Get(0).(func(context.Context) []*T); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*T)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: ctx, params
func (_m *ICRUDRepository[T]) Create(ctx context.Context, params *T) (*T, error) {
	ret := _m.Called(ctx, params)

	var r0 *T
	if rf, ok := ret.Get(0).(func(context.Context, *T) *T); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*T)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *T) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, params
func (_m *ICRUDRepository[T]) Delete(ctx context.Context, params *T) error {
	ret := _m.Called(ctx, params)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *T) error); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Find provides a mock function with given fields: ctx, key, val
func (_m *ICRUDRepository[T]) Find(ctx context.Context, key domain.FindWith, val interface{}) (*T, error) {
	ret := _m.Called(ctx, key, val)

	var r0 *T
	if rf, ok := ret.Get(0).(func(context.Context, domain.FindWith, interface{}) *T); ok {
		r0 = rf(ctx, key, val)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*T)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.FindWith, interface{}) error); ok {
		r1 = rf(ctx, key, val)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, params
func (_m *ICRUDRepository[T]) Update(ctx context.Context, params *T) (*T, error) {
	ret := _m.Called(ctx, params)

	var r0 *T
	if rf, ok := ret.Get(0).(func(context.Context, *T) *T); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*T)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *T) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewICRUDRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewICRUDRepository creates a new instance of ICRUDRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewICRUDRepository[T interface{}](t mockConstructorTestingTNewICRUDRepository) *ICRUDRepository[T] {
	mock := &ICRUDRepository[T]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
