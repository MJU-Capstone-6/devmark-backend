package interfaces

import "github.com/MJU-Capstone-6/devmark-backend/internal/repository"

//go:generate mockery --name IDeviceinfoService
type IDeviceinfoService interface {
	Create(repository.CreateDeviceInfoParams) (*repository.DeviceInfo, error)
	CheckDeviceInfoExists(string) error
}
