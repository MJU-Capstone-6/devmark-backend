// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	repository "github.com/MJU-Capstone-6/devmark-backend/internal/repository"
)

// IRepository is an autogenerated mock type for the IRepository type
type IRepository struct {
	mock.Mock
}

// CreateRefreshToken provides a mock function with given fields: _a0, _a1
func (_m *IRepository) CreateRefreshToken(_a0 context.Context, _a1 repository.CreateRefreshTokenParams) (repository.RefreshToken, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CreateRefreshToken")
	}

	var r0 repository.RefreshToken
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, repository.CreateRefreshTokenParams) (repository.RefreshToken, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, repository.CreateRefreshTokenParams) repository.RefreshToken); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(repository.RefreshToken)
	}

	if rf, ok := ret.Get(1).(func(context.Context, repository.CreateRefreshTokenParams) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindRefreshTokenByUserID provides a mock function with given fields: _a0, _a1
func (_m *IRepository) FindRefreshTokenByUserID(_a0 context.Context, _a1 *int32) (repository.RefreshToken, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for FindRefreshTokenByUserID")
	}

	var r0 repository.RefreshToken
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *int32) (repository.RefreshToken, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *int32) repository.RefreshToken); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(repository.RefreshToken)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *int32) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindUserByUsername provides a mock function with given fields: _a0, _a1
func (_m *IRepository) FindUserByUsername(_a0 context.Context, _a1 *string) (repository.User, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for FindUserByUsername")
	}

	var r0 repository.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *string) (repository.User, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *string) repository.User); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(repository.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateRefreshToken provides a mock function with given fields: _a0, _a1
func (_m *IRepository) UpdateRefreshToken(_a0 context.Context, _a1 repository.UpdateRefreshTokenParams) (repository.RefreshToken, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for UpdateRefreshToken")
	}

	var r0 repository.RefreshToken
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, repository.UpdateRefreshTokenParams) (repository.RefreshToken, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, repository.UpdateRefreshTokenParams) repository.RefreshToken); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(repository.RefreshToken)
	}

	if rf, ok := ret.Get(1).(func(context.Context, repository.UpdateRefreshTokenParams) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIRepository creates a new instance of IRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *IRepository {
	mock := &IRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
