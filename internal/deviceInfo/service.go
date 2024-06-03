package deviceinfo

import (
	"context"

	"github.com/MJU-Capstone-6/devmark-backend/internal/repository"
	"github.com/MJU-Capstone-6/devmark-backend/pkg/interfaces"
)

type DeviceInfoService struct {
	Repository interfaces.IRepository
}

func InitDeviceInfoService() *DeviceInfoService {
	return &DeviceInfoService{}
}

func (d *DeviceInfoService) Create(param repository.CreateDeviceInfoParams) (*repository.DeviceInfo, error) {
	deviceInfo, err := d.Repository.CreateDeviceInfo(context.Background(), param)
	if err != nil {
		return nil, err
	}
	return &deviceInfo, nil
}

func (d *DeviceInfoService) CheckDeviceInfoExists(token string) error {
	_, err := d.Repository.FindDeviceInfoByToken(context.Background(), &token)
	if err != nil {
		return err
	}
	return nil
}

func (d DeviceInfoService) WithRepository(repo interfaces.IRepository) DeviceInfoService {
	d.Repository = repo
	return d
}
