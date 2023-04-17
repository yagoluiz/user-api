// Code generated by mockery v2.25.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	domain "github.com/yagoluiz/user-api/internal/domain"
)

// UserSearchUseCaseInterface is an autogenerated mock type for the UserSearchUseCaseInterface type
type UserSearchUseCaseInterface struct {
	mock.Mock
}

// FindUser provides a mock function with given fields: term, limit, page
func (_m *UserSearchUseCaseInterface) FindUser(term string, limit int, page int) ([]*domain.User, error) {
	ret := _m.Called(term, limit, page)

	var r0 []*domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string, int, int) ([]*domain.User, error)); ok {
		return rf(term, limit, page)
	}
	if rf, ok := ret.Get(0).(func(string, int, int) []*domain.User); ok {
		r0 = rf(term, limit, page)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string, int, int) error); ok {
		r1 = rf(term, limit, page)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserSearchUseCaseInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserSearchUseCaseInterface creates a new instance of UserSearchUseCaseInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserSearchUseCaseInterface(t mockConstructorTestingTNewUserSearchUseCaseInterface) *UserSearchUseCaseInterface {
	mock := &UserSearchUseCaseInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
