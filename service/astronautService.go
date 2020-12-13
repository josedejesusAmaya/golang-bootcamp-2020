//go:generate mockgen -destination=mocks/service/mockAstronautService.go -package=service . AstronautService
package service

import (
	"github.com/josedejesusAmaya/golang-bootcamp-2020/domain"
	"github.com/josedejesusAmaya/golang-bootcamp-2020/errs"
)

// AstronautService is the port for the User adapter
type AstronautService interface {
	GetAllAstronauts() ([]domain.Astronaut, *errs.AppError)
	GetAscAstronauts() ([]domain.Astronaut, *errs.AppError)
	GetDescAstronauts() ([]domain.Astronaut, *errs.AppError)
}

// DefaultAstronautService defines the type of data for the Data interface
type DefaultAstronautService struct {
	repository domain.AstronautRepository
}

// GetAllAstronauts is the implementation to store the data
func (d DefaultAstronautService) GetAllAstronauts() ([]domain.Astronaut, *errs.AppError) {
	return d.repository.FindAll()
}

// GetAscAstronauts is the implementation to store the asc data
func (d DefaultAstronautService) GetAscAstronauts() ([]domain.Astronaut, *errs.AppError) {
	return d.repository.FindAsc()
}

// GetDescAstronauts is the implementation to store the desc data
func (d DefaultAstronautService) GetDescAstronauts() ([]domain.Astronaut, *errs.AppError) {
	return d.repository.FindDesc()
}

// NewAstronautService to initialize the service
func NewAstronautService(a domain.AstronautRepository) DefaultAstronautService {
	return DefaultAstronautService{repository: a}
}
