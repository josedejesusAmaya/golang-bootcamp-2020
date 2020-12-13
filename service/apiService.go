//go:generate mockgen -destination=mocks/service/mockAPIService.go -package=service . APIService
package service

import (
	"github.com/josedejesusAmaya/golang-bootcamp-2020/domain"
	"github.com/josedejesusAmaya/golang-bootcamp-2020/errs"
)

// APIService is the port for the API adapter
type APIService interface {
	WriteCSV() (string, *errs.AppError)
}

// DefaultAPIService defines the type of data for the API interface
type DefaultAPIService struct {
	repository domain.APIRepository
}

// WriteCSV is the implementation to make the HTTP request, create and write the file
func (api DefaultAPIService) WriteCSV() (string, *errs.AppError) {
	return api.repository.FindAll()
}

// NewAPIService to initialize the service
func NewAPIService(api domain.APIRepository) DefaultAPIService {
	return DefaultAPIService{repository: api}
}
