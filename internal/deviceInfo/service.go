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

func (d *DeviceInfoService) FindByUserID(id int) (*repository.DeviceInfo, error) {
	deviceInfo, err := d.Repository.FindDeviceInfo(context.Background(), int64(id))
	if err != nil {
		return nil, err
	}
	return &deviceInfo, nil
}

func (d *DeviceInfoService) FindByAgentAndUserID(id int, agent string) (*repository.DeviceInfo, error) {
	param := repository.FindDeviceInfoByAgentAndUserIDParams{
		UserID:      int64(id),
		AgentHeader: agent,
	}
	deviceInfo, err := d.Repository.FindDeviceInfoByAgentAndUserID(context.Background(), param)
	if err != nil {
		return nil, err
	}
	return &deviceInfo, nil
}

func (d *DeviceInfoService) FindByAgent(agent string) (*repository.DeviceInfo, error) {
	deviceInfo, err := d.Repository.FindDeviceInfoByAgent(context.Background(), agent)
	if err != nil {
		return nil, err
	}
	return &deviceInfo, nil
}

func (d DeviceInfoService) WithRepository(repo interfaces.IRepository) DeviceInfoService {
	d.Repository = repo
	return d
}
