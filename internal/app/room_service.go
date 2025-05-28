package app

import (
	"log"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
)

type RoomService interface {
	Save(h domain.Room) (domain.Room, error)
	Find(id uint64) (interface{}, error)
	FindByHouseId(hid uint64) ([]domain.Room, error)
	Update(r, newR domain.Room) (domain.Room, error)
	Delete(id uint64) error
}

type roomService struct {
	roomRepo database.RoomRepository
}

func NewRoomService(rr database.RoomRepository) RoomService {
	return roomService{
		roomRepo: rr,
	}
}

func (s roomService) Save(r domain.Room) (domain.Room, error) {
	room, err := s.roomRepo.Save(r)
	if err != nil {
		log.Printf("roomService.Save(s.roomRepo.Save): %s", err)
		return domain.Room{}, err
	}

	return room, nil
}

func (s roomService) Find(id uint64) (interface{}, error) {
	room, err := s.roomRepo.Find(id)
	if err != nil {
		log.Printf("roomService.Find(s.roomRepo.Find): %s", err)
		return domain.Room{}, err
	}
	return room, nil
}

func (s roomService) FindByHouseId(hid uint64) ([]domain.Room, error) {
	rooms, err := s.roomRepo.FindByHouseId(hid)
	if err != nil {
		log.Printf("roomService.FindByHouseId(s.roomRepo.FindByHouseId): %s", err)
		return nil, err
	}

	return rooms, nil
}

func (s roomService) Update(r, newR domain.Room) (domain.Room, error) {
	if newR.Name != "" {
		r.Name = newR.Name
	}
	if newR.Description != nil {
		r.Description = newR.Description
	}

	room, err := s.roomRepo.Update(r)
	if err != nil {
		log.Printf("roomService.Update(s.roomRepo.Update): %s", err)
		return domain.Room{}, err
	}
	return room, nil
}

func (s roomService) Delete(id uint64) error {
	err := s.roomRepo.Delete(id)
	if err != nil {
		log.Printf("roomService.Delete(s.roomRepo.Delete): %s", err)
		return err
	}
	return nil
}
