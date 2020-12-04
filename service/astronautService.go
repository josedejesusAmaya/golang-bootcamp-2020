package service

import "github.com/josedejesusAmaya/golang-bootcamp-2020/domain"

// AstronautService is the port for the User side
type AstronautService interface {
	GetAllAstronauts() ([]domain.Astronaut, error)
}

// DefaultAstronautService is
type DefaultAstronautService struct {
	repository domain.AstronautRepository
}

// GetAllAstronauts is
func (d DefaultAstronautService) GetAllAstronauts() ([]domain.Astronaut, error) {
	return d.repository.FindAll()
}

// NewAstronautService is
func NewAstronautService(a domain.AstronautRepository) DefaultAstronautService {
	return DefaultAstronautService{repository: a}
}
