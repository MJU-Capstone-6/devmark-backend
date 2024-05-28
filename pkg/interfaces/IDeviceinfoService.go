package interfaces

import "github.com/MJU-Capstone-6/devmark-backend/internal/repository"

//go:generate mockery --name IDeviceinfoService
type IDeviceinfoService interface {
	Create(repository.CreateDeviceInfoParams) (*repository.DeviceInfo, error)
	FindByUserID(int) (*repository.DeviceInfo, error)
	FindByAgent(string) (*repository.DeviceInfo, error)
	FindByAgentAndUserID(int, string) (*repository.DeviceInfo, error)
}
