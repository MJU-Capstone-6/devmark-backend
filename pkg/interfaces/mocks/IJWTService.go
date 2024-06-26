// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	paseto "github.com/o1egl/paseto"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// IJWTService is an autogenerated mock type for the IJWTService type
type IJWTService struct {
	mock.Mock
}

// GenerateToken provides a mock function with given fields: _a0, _a1
func (_m *IJWTService) GenerateToken(_a0 int, _a1 time.Time) (string, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GenerateToken")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(int, time.Time) (string, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(int, time.Time) string); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(int, time.Time) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VerifyToken provides a mock function with given fields: _a0
func (_m *IJWTService) VerifyToken(_a0 string) (paseto.JSONToken, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for VerifyToken")
	}

	var r0 paseto.JSONToken
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (paseto.JSONToken, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(string) paseto.JSONToken); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(paseto.JSONToken)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIJWTService creates a new instance of IJWTService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIJWTService(t interface {
	mock.TestingT
	Cleanup(func())
}) *IJWTService {
	mock := &IJWTService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
