// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	repository "github.com/MJU-Capstone-6/devmark-backend/internal/repository"
	mock "github.com/stretchr/testify/mock"
)

// IWorkspaceCodeService is an autogenerated mock type for the IWorkspaceCodeService type
type IWorkspaceCodeService struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0
func (_m *IWorkspaceCodeService) Create(_a0 repository.CreateWorkspaceCodeParams) (*repository.WorkspaceCode, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *repository.WorkspaceCode
	var r1 error
	if rf, ok := ret.Get(0).(func(repository.CreateWorkspaceCodeParams) (*repository.WorkspaceCode, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(repository.CreateWorkspaceCodeParams) *repository.WorkspaceCode); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repository.WorkspaceCode)
		}
	}

	if rf, ok := ret.Get(1).(func(repository.CreateWorkspaceCodeParams) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateCategoryAndBookmark provides a mock function with given fields: _a0
func (_m *IWorkspaceCodeService) CreateCategoryAndBookmark(_a0 string) (*repository.Bookmark, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for CreateCategoryAndBookmark")
	}

	var r0 *repository.Bookmark
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*repository.Bookmark, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(string) *repository.Bookmark); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repository.Bookmark)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByCode provides a mock function with given fields: _a0
func (_m *IWorkspaceCodeService) FindByCode(_a0 string) (*repository.WorkspaceCode, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for FindByCode")
	}

	var r0 *repository.WorkspaceCode
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*repository.WorkspaceCode, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(string) *repository.WorkspaceCode); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repository.WorkspaceCode)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: _a0
func (_m *IWorkspaceCodeService) Update(_a0 repository.UpdateWorkspaceCodeParams) (*repository.WorkspaceCode, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *repository.WorkspaceCode
	var r1 error
	if rf, ok := ret.Get(0).(func(repository.UpdateWorkspaceCodeParams) (*repository.WorkspaceCode, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(repository.UpdateWorkspaceCodeParams) *repository.WorkspaceCode); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repository.WorkspaceCode)
		}
	}

	if rf, ok := ret.Get(1).(func(repository.UpdateWorkspaceCodeParams) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIWorkspaceCodeService creates a new instance of IWorkspaceCodeService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIWorkspaceCodeService(t interface {
	mock.TestingT
	Cleanup(func())
}) *IWorkspaceCodeService {
	mock := &IWorkspaceCodeService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}