package requests

import "github.com/BohdanBoriak/boilerplate-go-back/internal/domain"

type HouseRequest struct {
	Name        string  `json:"name" validate:"required"`
	Description *string `json:"description"`
	City        string  `json:"city" validate:"required"`
	Address     string  `json:"address" validate:"required"`
	Lat         float64 `json:"lat" validate:"required"`
	Lon         float64 `json:"lon" validate:"required"`
}

type UpdateHouseRequest struct {
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	City        *string  `json:"city"`
	Address     *string  `json:"address"`
	Lat         *float64 `json:"lat"`
	Lon         *float64 `json:"lon"`
}

func (r HouseRequest) ToDomainModel() (interface{}, error) {
	return domain.House{
		Name:        r.Name,
		Description: r.Description,
		City:        r.City,
		Address:     r.Address,
		Lat:         r.Lat,
		Lon:         r.Lon,
	}, nil
}

func (r UpdateHouseRequest) ToDomainModel() (interface{}, error) {
	var (
		name    string
		city    string
		address string
		lat     float64
		lon     float64
	)
	if r.Name != nil {
		name = *r.Name
	}
	if r.City != nil {
		city = *r.City
	}
	if r.Address != nil {
		address = *r.Address
	}
	if r.Lat != nil {
		lat = *r.Lat
	}
	if r.Lon != nil {
		lon = *r.Lon
	}

	return domain.House{
		Name:        name,
		Description: r.Description,
		City:        city,
		Address:     address,
		Lat:         lat,
		Lon:         lon,
	}, nil
}
