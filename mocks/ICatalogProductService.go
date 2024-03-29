// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/aasumitro/posbe/domain"
	mock "github.com/stretchr/testify/mock"

	utils "github.com/aasumitro/posbe/pkg/utils"
)

// ICatalogProductService is an autogenerated mock type for the ICatalogProductService type
type ICatalogProductService struct {
	mock.Mock
}

// AddProduct provides a mock function with given fields: data
func (_m *ICatalogProductService) AddProduct(data *domain.Product) (*domain.Product, *utils.ServiceError) {
	ret := _m.Called(data)

	var r0 *domain.Product
	if rf, ok := ret.Get(0).(func(*domain.Product) *domain.Product); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Product)
		}
	}

	var r1 *utils.ServiceError
	if rf, ok := ret.Get(1).(func(*domain.Product) *utils.ServiceError); ok {
		r1 = rf(data)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*utils.ServiceError)
		}
	}

	return r0, r1
}

// AddProductVariant provides a mock function with given fields: data
func (_m *ICatalogProductService) AddProductVariant(data *domain.ProductVariant) (*domain.ProductVariant, *utils.ServiceError) {
	ret := _m.Called(data)

	var r0 *domain.ProductVariant
	if rf, ok := ret.Get(0).(func(*domain.ProductVariant) *domain.ProductVariant); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.ProductVariant)
		}
	}

	var r1 *utils.ServiceError
	if rf, ok := ret.Get(1).(func(*domain.ProductVariant) *utils.ServiceError); ok {
		r1 = rf(data)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*utils.ServiceError)
		}
	}

	return r0, r1
}

// DeleteProduct provides a mock function with given fields: data
func (_m *ICatalogProductService) DeleteProduct(data *domain.Product) *utils.ServiceError {
	ret := _m.Called(data)

	var r0 *utils.ServiceError
	if rf, ok := ret.Get(0).(func(*domain.Product) *utils.ServiceError); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*utils.ServiceError)
		}
	}

	return r0
}

// DeleteProductVariant provides a mock function with given fields: data
func (_m *ICatalogProductService) DeleteProductVariant(data *domain.ProductVariant) *utils.ServiceError {
	ret := _m.Called(data)

	var r0 *utils.ServiceError
	if rf, ok := ret.Get(0).(func(*domain.ProductVariant) *utils.ServiceError); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*utils.ServiceError)
		}
	}

	return r0
}

// EditProduct provides a mock function with given fields: data
func (_m *ICatalogProductService) EditProduct(data *domain.Product) (*domain.Product, *utils.ServiceError) {
	ret := _m.Called(data)

	var r0 *domain.Product
	if rf, ok := ret.Get(0).(func(*domain.Product) *domain.Product); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Product)
		}
	}

	var r1 *utils.ServiceError
	if rf, ok := ret.Get(1).(func(*domain.Product) *utils.ServiceError); ok {
		r1 = rf(data)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*utils.ServiceError)
		}
	}

	return r0, r1
}

// EditProductVariant provides a mock function with given fields: data
func (_m *ICatalogProductService) EditProductVariant(data *domain.ProductVariant) (*domain.ProductVariant, *utils.ServiceError) {
	ret := _m.Called(data)

	var r0 *domain.ProductVariant
	if rf, ok := ret.Get(0).(func(*domain.ProductVariant) *domain.ProductVariant); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.ProductVariant)
		}
	}

	var r1 *utils.ServiceError
	if rf, ok := ret.Get(1).(func(*domain.ProductVariant) *utils.ServiceError); ok {
		r1 = rf(data)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*utils.ServiceError)
		}
	}

	return r0, r1
}

// ProductDetail provides a mock function with given fields: id
func (_m *ICatalogProductService) ProductDetail(id int) (*domain.Product, *utils.ServiceError) {
	ret := _m.Called(id)

	var r0 *domain.Product
	if rf, ok := ret.Get(0).(func(int) *domain.Product); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Product)
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

// ProductList provides a mock function with given fields:
func (_m *ICatalogProductService) ProductList() ([]*domain.Product, *utils.ServiceError) {
	ret := _m.Called()

	var r0 []*domain.Product
	if rf, ok := ret.Get(0).(func() []*domain.Product); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Product)
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

// ProductSearch provides a mock function with given fields: keys, values
func (_m *ICatalogProductService) ProductSearch(keys []domain.FindWith, values []interface{}) ([]*domain.Product, *utils.ServiceError) {
	ret := _m.Called(keys, values)

	var r0 []*domain.Product
	if rf, ok := ret.Get(0).(func([]domain.FindWith, []interface{}) []*domain.Product); ok {
		r0 = rf(keys, values)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Product)
		}
	}

	var r1 *utils.ServiceError
	if rf, ok := ret.Get(1).(func([]domain.FindWith, []interface{}) *utils.ServiceError); ok {
		r1 = rf(keys, values)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*utils.ServiceError)
		}
	}

	return r0, r1
}

type mockConstructorTestingTNewICatalogProductService interface {
	mock.TestingT
	Cleanup(func())
}

// NewICatalogProductService creates a new instance of ICatalogProductService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewICatalogProductService(t mockConstructorTestingTNewICatalogProductService) *ICatalogProductService {
	mock := &ICatalogProductService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
