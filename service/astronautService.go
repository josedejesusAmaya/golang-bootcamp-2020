package service

import "github.com/josedejesusAmaya/golang-bootcamp-2020/domain"

// AstronautService is the port for the User adapter
type AstronautService interface {
	GetAllAstronauts() ([]domain.Astronaut, error)
}

// DefaultAstronautService defines the type of data for the Data interface
type DefaultAstronautService struct {
	repository domain.AstronautRepository
}

// GetAllAstronauts is the implementation to store the data
func (d DefaultAstronautService) GetAllAstronauts() ([]domain.Astronaut, error) {
	return d.repository.FindAll()
}

// NewAstronautService to initialize the service
func NewAstronautService(a domain.AstronautRepository) DefaultAstronautService {
	return DefaultAstronautService{repository: a}
}
