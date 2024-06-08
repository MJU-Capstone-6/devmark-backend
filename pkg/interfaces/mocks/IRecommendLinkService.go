// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	repository "github.com/MJU-Capstone-6/devmark-backend/internal/repository"
	mock "github.com/stretchr/testify/mock"
)

// IRecommendLinkService is an autogenerated mock type for the IRecommendLinkService type
type IRecommendLinkService struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0
func (_m *IRecommendLinkService) Create(_a0 repository.CreateRecommendLinkParams) (*repository.RecommendLink, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *repository.RecommendLink
	var r1 error
	if rf, ok := ret.Get(0).(func(repository.CreateRecommendLinkParams) (*repository.RecommendLink, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(repository.CreateRecommendLinkParams) *repository.RecommendLink); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repository.RecommendLink)
		}
	}

	if rf, ok := ret.Get(1).(func(repository.CreateRecommendLinkParams) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIRecommendLinkService creates a new instance of IRecommendLinkService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIRecommendLinkService(t interface {
	mock.TestingT
	Cleanup(func())
}) *IRecommendLinkService {
	mock := &IRecommendLinkService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}