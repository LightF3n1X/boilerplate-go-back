package app

import "github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"

type MeasurementsService interface {
}

type deviceMeasurementsService struct {
	measurementsRepo database.MeasurementsRepository
}

func NewMeasurementsService(mr database.MeasurementsRepository) MeasurementsService {
	return deviceMeasurementsService{
		measurementsRepo: mr,
	}
}
