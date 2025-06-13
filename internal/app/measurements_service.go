package app

import (
	"log"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
)

type MeasurementsService interface {
	Save(m domain.Measurement) (domain.Measurement, error)
}

type measurementsService struct {
	measurementsRepo database.MeasurementsRepository
}

func NewMeasurementsService(mr database.MeasurementsRepository) MeasurementsService {
	return measurementsService{
		measurementsRepo: mr,
	}
}

func (s measurementsService) Save(m domain.Measurement) (domain.Measurement, error) {
	measurements, err := s.measurementsRepo.Save(m)
	if err != nil {
		log.Printf("measurementsService.Save(s.measurementsRepo.Save): %s", err)
		return domain.Measurement{}, err
	}

	return measurements, nil
}
