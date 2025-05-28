package app

import (
	"log"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
)

type DeviceService interface {
	Save(d domain.Device) (domain.Device, error)
	Find(id uint64) (interface{}, error)
}

type deviceService struct {
	deviceRepo database.DeviceRepository
}

func NewDeviceService(dr database.DeviceRepository) DeviceService {
	return deviceService{
		deviceRepo: dr,
	}
}

func (s deviceService) Save(d domain.Device) (domain.Device, error) {
	device, err := s.deviceRepo.Save(d)
	if err != nil {
		log.Printf("deviceService.Save(s.deviceRepo.Save): %s", err)
		return domain.Device{}, err
	}

	return device, nil
}

func (s deviceService) Find(id uint64) (interface{}, error) {
	device, err := s.deviceRepo.Find(id)
	if err != nil {
		log.Printf("deviceService.Find(s.deviceRepo.Find): %s", err)
		return domain.Device{}, err
	}
	return device, nil
}
