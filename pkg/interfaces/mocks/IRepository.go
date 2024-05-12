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

// CreateInviteCode provides a mock function with given fields: _a0, _a1
func (_m *IRepository) CreateInviteCode(_a0 context.Context, _a1 repository.CreateInviteCodeParams) (repository.InviteCode, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CreateInviteCode")
	}

	var r0 repository.InviteCode
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, repository.CreateInviteCodeParams) (repository.InviteCode, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, repository.CreateInviteCodeParams) repository.InviteCode); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(repository.InviteCode)
	}

	if rf, ok := ret.Get(1).(func(context.Context, repository.CreateInviteCodeParams) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
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

// CreateUser provides a mock function with given fields: _a0, _a1
func (_m *IRepository) CreateUser(_a0 context.Context, _a1 repository.CreateUserParams) (repository.User, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CreateUser")
	}

	var r0 repository.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, repository.CreateUserParams) (repository.User, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, repository.CreateUserParams) repository.User); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(repository.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, repository.CreateUserParams) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateWorkspace provides a mock function with given fields: _a0, _a1
func (_m *IRepository) CreateWorkspace(_a0 context.Context, _a1 *string) (repository.Workspace, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CreateWorkspace")
	}

	var r0 repository.Workspace
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *string) (repository.Workspace, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *string) repository.Workspace); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(repository.Workspace)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteWorkspace provides a mock function with given fields: _a0, _a1
func (_m *IRepository) DeleteWorkspace(_a0 context.Context, _a1 int64) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for DeleteWorkspace")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindInviteCodeByWorkspaceID provides a mock function with given fields: _a0, _a1
func (_m *IRepository) FindInviteCodeByWorkspaceID(_a0 context.Context, _a1 *int32) (repository.InviteCode, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for FindInviteCodeByWorkspaceID")
	}

	var r0 repository.InviteCode
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *int32) (repository.InviteCode, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *int32) repository.InviteCode); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(repository.InviteCode)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *int32) error); ok {
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

// FindWorkspace provides a mock function with given fields: _a0, _a1
func (_m *IRepository) FindWorkspace(_a0 context.Context, _a1 int64) (repository.WorkspaceUserCategory, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for FindWorkspace")
	}

	var r0 repository.WorkspaceUserCategory
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (repository.WorkspaceUserCategory, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) repository.WorkspaceUserCategory); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(repository.WorkspaceUserCategory)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// JoinWorkspace provides a mock function with given fields: _a0, _a1
func (_m *IRepository) JoinWorkspace(_a0 context.Context, _a1 repository.JoinWorkspaceParams) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for JoinWorkspace")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, repository.JoinWorkspaceParams) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
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

// UpdateUser provides a mock function with given fields: _a0, _a1
func (_m *IRepository) UpdateUser(_a0 context.Context, _a1 repository.UpdateUserParams) (repository.User, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUser")
	}

	var r0 repository.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, repository.UpdateUserParams) (repository.User, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, repository.UpdateUserParams) repository.User); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(repository.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, repository.UpdateUserParams) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateWorkspace provides a mock function with given fields: _a0, _a1
func (_m *IRepository) UpdateWorkspace(_a0 context.Context, _a1 repository.UpdateWorkspaceParams) (repository.Workspace, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for UpdateWorkspace")
	}

	var r0 repository.Workspace
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, repository.UpdateWorkspaceParams) (repository.Workspace, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, repository.UpdateWorkspaceParams) repository.Workspace); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(repository.Workspace)
	}

	if rf, ok := ret.Get(1).(func(context.Context, repository.UpdateWorkspaceParams) error); ok {
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
