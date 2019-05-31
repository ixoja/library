// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import models "github.com/ixoja/library/internal/models"

// Storage is an autogenerated mock type for the Storage type
type Storage struct {
	mock.Mock
}

// Create provides a mock function with given fields: book
func (_m *Storage) Create(book *models.Book) (*models.Book, error) {
	ret := _m.Called(book)

	var r0 *models.Book
	if rf, ok := ret.Get(0).(func(*models.Book) *models.Book); ok {
		r0 = rf(book)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Book)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.Book) error); ok {
		r1 = rf(book)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *Storage) Delete(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: id
func (_m *Storage) Get(id string) (*models.Book, bool, error) {
	ret := _m.Called(id)

	var r0 *models.Book
	if rf, ok := ret.Get(0).(func(string) *models.Book); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Book)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(id)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetAll provides a mock function with given fields:
func (_m *Storage) GetAll() ([]*models.Book, error) {
	ret := _m.Called()

	var r0 []*models.Book
	if rf, ok := ret.Get(0).(func() []*models.Book); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Book)
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

// Update provides a mock function with given fields: book
func (_m *Storage) Update(book *models.Book) error {
	ret := _m.Called(book)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Book) error); ok {
		r0 = rf(book)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
