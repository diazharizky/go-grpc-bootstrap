// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	pb "github.com/diazharizky/go-grpc-bootstrap/pb"
	mock "github.com/stretchr/testify/mock"
)

// IUserRepository is an autogenerated mock type for the IUserRepository type
type IUserRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: params
func (_m *IUserRepository) Create(params pb.User) error {
	ret := _m.Called(params)

	var r0 error
	if rf, ok := ret.Get(0).(func(pb.User) error); ok {
		r0 = rf(params)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: username
func (_m *IUserRepository) Delete(username string) error {
	ret := _m.Called(username)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields:
func (_m *IUserRepository) Get() ([]*pb.User, error) {
	ret := _m.Called()

	var r0 []*pb.User
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*pb.User, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*pb.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*pb.User)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields:
func (_m *IUserRepository) List() ([]*pb.User, error) {
	ret := _m.Called()

	var r0 []*pb.User
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*pb.User, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*pb.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*pb.User)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: params
func (_m *IUserRepository) Update(params pb.User) error {
	ret := _m.Called(params)

	var r0 error
	if rf, ok := ret.Get(0).(func(pb.User) error); ok {
		r0 = rf(params)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewIUserRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewIUserRepository creates a new instance of IUserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIUserRepository(t mockConstructorTestingTNewIUserRepository) *IUserRepository {
	mock := &IUserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
