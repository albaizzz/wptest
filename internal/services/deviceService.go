package services

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/wptest/internal/models"
	"github.com/wptest/internal/repositories"
)

type DeviceService struct {
	DeviceRepository  repositories.IDeviceRepository
	MessageRepository repositories.IMessageRepository
}

type IDeviceService interface {
	Publish(device models.Device) error
	GetById(id uint64) (models.Device, error)
	GetAll() ([]models.Device, error)
	ReceiveDevice(msg []byte) (err error)
}

func NewDeviceService(deviceRepository repositories.IDeviceRepository, messageRepository repositories.IMessageRepository) *DeviceService {
	return &DeviceService{
		DeviceRepository:  deviceRepository,
		MessageRepository: messageRepository,
	}
}

func (ds *DeviceService) GetById(id uint64) (device models.Device, err error) {
	device, err = ds.DeviceRepository.GetById(id)
	return
}

func (ds *DeviceService) GetAll() (devices []models.Device, err error) {
	devices, err = ds.DeviceRepository.GetAll()
	return
}

func (ds *DeviceService) Publish(device models.Device) error {
	err := ds.MessageRepository.Publish(device)
	return err
}

func (ds *DeviceService) ReceiveDevice(msg []byte) (err error) {
	fmt.Println(string(msg))
	var devices models.Device
	err = json.Unmarshal(msg, &devices)
	if err != nil {
		return err
	}
	devices.UpdatedAt = time.Now()
	err = ds.DeviceRepository.Store(devices)
	if err != nil {
		return err
	}
	return nil
}
