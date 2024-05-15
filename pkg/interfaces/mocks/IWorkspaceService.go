// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	repository "github.com/MJU-Capstone-6/devmark-backend/internal/repository"
	mock "github.com/stretchr/testify/mock"
)

// IWorkspaceService is an autogenerated mock type for the IWorkspaceService type
type IWorkspaceService struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0, _a1
func (_m *IWorkspaceService) Create(_a0 int, _a1 repository.CreateWorkspaceParams) (*repository.Workspace, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *repository.Workspace
	var r1 error
	if rf, ok := ret.Get(0).(func(int, repository.CreateWorkspaceParams) (*repository.Workspace, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(int, repository.CreateWorkspaceParams) *repository.Workspace); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repository.Workspace)
		}
	}

	if rf, ok := ret.Get(1).(func(int, repository.CreateWorkspaceParams) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: _a0
func (_m *IWorkspaceService) Delete(_a0 int) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindById provides a mock function with given fields: _a0
func (_m *IWorkspaceService) FindById(_a0 int) (*repository.WorkspaceUserCategory, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for FindById")
	}

	var r0 *repository.WorkspaceUserCategory
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*repository.WorkspaceUserCategory, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(int) *repository.WorkspaceUserCategory); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repository.WorkspaceUserCategory)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindCategoriesById provides a mock function with given fields: _a0
func (_m *IWorkspaceService) FindCategoriesById(_a0 int) (*[]repository.Category, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for FindCategoriesById")
	}

	var r0 *[]repository.Category
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*[]repository.Category, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(int) *[]repository.Category); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]repository.Category)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Join provides a mock function with given fields: _a0, _a1
func (_m *IWorkspaceService) Join(_a0 string, _a1 repository.JoinWorkspaceParams) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Join")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, repository.JoinWorkspaceParams) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: _a0
func (_m *IWorkspaceService) Update(_a0 repository.UpdateWorkspaceParams) (*repository.Workspace, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *repository.Workspace
	var r1 error
	if rf, ok := ret.Get(0).(func(repository.UpdateWorkspaceParams) (*repository.Workspace, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(repository.UpdateWorkspaceParams) *repository.Workspace); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repository.Workspace)
		}
	}

	if rf, ok := ret.Get(1).(func(repository.UpdateWorkspaceParams) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIWorkspaceService creates a new instance of IWorkspaceService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIWorkspaceService(t interface {
	mock.TestingT
	Cleanup(func())
}) *IWorkspaceService {
	mock := &IWorkspaceService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
