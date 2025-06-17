package app

import (
	"log"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
)

type DeviceService interface {
	Save(d domain.Device) (domain.Device, error)
	Find(id uint64) (interface{}, error)
	FindById(id uint64) (domain.Device, error)
	Update(d, newD domain.Device) (domain.Device, error)
	Delete(id uint64) error
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

func (s deviceService) Update(d, newD domain.Device) (domain.Device, error) {
	if newD.SerialNumber != "" {
		d.SerialNumber = newD.SerialNumber
	}
	if newD.Characteristics != nil {
		d.Characteristics = newD.Characteristics
	}
	if newD.Category != "" {
		d.Category = newD.Category
	}
	if newD.Units != nil {
		d.Units = newD.Units
	}
	if newD.PowerConsumption != nil {
		d.PowerConsumption = newD.PowerConsumption
	}

	device, err := s.deviceRepo.Update(d)
	if err != nil {
		log.Printf("deviceService.Update(s.deviceRepo.Update): %s", err)
		return domain.Device{}, err
	}
	return device, nil
}

func (s deviceService) Delete(id uint64) error {
	err := s.deviceRepo.Delete(id)
	if err != nil {
		log.Printf("deviceService.Delete(s.deviceRepo.Delete): %s", err)
		return err
	}
	return nil
}

func (s deviceService) FindById(id uint64) (domain.Device, error) {
	device, err := s.deviceRepo.Find(id)
	if err != nil {
		log.Printf("deviceService.FindById(s.deviceRepo.Find): %s", err)
		return domain.Device{}, err
	}
	return device, nil
}
