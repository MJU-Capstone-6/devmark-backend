// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	repository "github.com/MJU-Capstone-6/devmark-backend/internal/repository"
	mock "github.com/stretchr/testify/mock"
)

// IDeviceinfoService is an autogenerated mock type for the IDeviceinfoService type
type IDeviceinfoService struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0
func (_m *IDeviceinfoService) Create(_a0 repository.CreateDeviceInfoParams) (*repository.DeviceInfo, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *repository.DeviceInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(repository.CreateDeviceInfoParams) (*repository.DeviceInfo, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(repository.CreateDeviceInfoParams) *repository.DeviceInfo); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repository.DeviceInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(repository.CreateDeviceInfoParams) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByAgent provides a mock function with given fields: _a0
func (_m *IDeviceinfoService) FindByAgent(_a0 string) (*repository.DeviceInfo, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for FindByAgent")
	}

	var r0 *repository.DeviceInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*repository.DeviceInfo, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(string) *repository.DeviceInfo); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repository.DeviceInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByUserID provides a mock function with given fields: _a0
func (_m *IDeviceinfoService) FindByUserID(_a0 int) (*repository.DeviceInfo, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for FindByUserID")
	}

	var r0 *repository.DeviceInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*repository.DeviceInfo, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(int) *repository.DeviceInfo); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repository.DeviceInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIDeviceinfoService creates a new instance of IDeviceinfoService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIDeviceinfoService(t interface {
	mock.TestingT
	Cleanup(func())
}) *IDeviceinfoService {
	mock := &IDeviceinfoService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
