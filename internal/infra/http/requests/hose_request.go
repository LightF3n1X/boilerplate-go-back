package requests

import "github.com/BohdanBoriak/boilerplate-go-back/internal/domain"

type HouseRequest struct {
	Name        string  `json:"name" validate:"required"`
	Description *string `json:"description"`
	City        string  `json:"city" validate:"required"`
	Adress      string  `json:"adress" validate:"required"`
	Lat         float64 `json:"lat" validate:"required"`
	Lon         float64 `json:"lon" validate:"required"`
}

func (r HouseRequest) ToDomainModel() (interface{}, error) {
	return domain.House{
		Name:        r.Name,
		Description: r.Description,
		City:        r.City,
		Adress:      r.Adress,
		Lat:         r.Lat,
		Lon:         r.Lon,
	}, nil
}
