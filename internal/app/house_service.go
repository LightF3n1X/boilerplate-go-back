package app

import (
	"log"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
)

type HouseService interface {
	Save(h domain.House) (domain.House, error)
	Find(id uint64) (interface{}, error)
	FindById(id uint64) (domain.House, error)
	FindList(uId uint64) ([]domain.House, error)
	Update(h, newH domain.House) (domain.House, error)
	Delete(id uint64) error
}

type houseService struct {
	houseRepo database.HouseRepository
	roomRepo  database.RoomRepository
}

func NewHouseService(hr database.HouseRepository, rr database.RoomRepository) HouseService {
	return houseService{
		houseRepo: hr,
		roomRepo:  rr,
	}
}

func (s houseService) Save(h domain.House) (domain.House, error) {
	house, err := s.houseRepo.Save(h)
	if err != nil {
		log.Printf("houseService.Save(s.houseRepo.Save): %s", err)
		return domain.House{}, err
	}

	return house, nil
}

func (s houseService) Find(id uint64) (interface{}, error) {
	house, err := s.houseRepo.Find(id)
	if err != nil {
		log.Printf("houseService.Find(s.houseRepo.Find): %s", err)
		return domain.House{}, err
	}

	rooms, err := s.roomRepo.FindByHouseId(id)
	if err != nil {
		log.Printf("houseService.Find(s.roomRepo.FindByHouseId): %s", err)
		return domain.House{}, err
	}
	house.Rooms = rooms
	return house, nil
}

func (s houseService) FindById(id uint64) (domain.House, error) {
	house, err := s.houseRepo.Find(id)
	if err != nil {
		log.Printf("houseService.Find(s.houseRepo.Find): %s", err)
		return domain.House{}, err
	}

	return house, nil
}

func (s houseService) FindList(uId uint64) ([]domain.House, error) {
	house, err := s.houseRepo.FindList(uId)
	if err != nil {
		log.Printf("houseService.FindList(s.houseRepo.FindList) :%s", err)
		return nil, err
	}

	return house, nil
}

func (s houseService) Update(h, newH domain.House) (domain.House, error) {
	if newH.Name != "" {
		h.Name = newH.Name
	}
	if newH.Description != nil {
		h.Description = newH.Description
	}
	if newH.City != "" {
		h.City = newH.City
	}
	if newH.Address != "" {
		h.Address = newH.Address
	}
	if newH.Lat != 0 {
		h.Lat = newH.Lat
	}
	if newH.Lon != 0 {
		h.Lon = newH.Lon
	}

	house, err := s.houseRepo.Update(h)
	if err != nil {
		log.Printf("houseService.Update(s.houseRepo.Update): %s", err)
		return domain.House{}, err
	}

	return house, nil
}

func (s houseService) Delete(id uint64) error {
	err := s.houseRepo.Delete(id)
	if err != nil {
		log.Printf("houseService.Delete(s.houseRepo.Delete): %s", err)
		return err
	}
	return nil
}
