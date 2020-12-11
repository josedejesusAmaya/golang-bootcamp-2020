package service

import "github.com/josedejesusAmaya/golang-bootcamp-2020/domain"

// APIService is the port for the API adapter 
type APIService interface {
	WriteCSV() (string, error)
}

// DefaultAPIService defines the type of data for the API interface
type DefaultAPIService struct {
	repository domain.APIRepository
}

// WriteCSV is the implementation to make the HTTP request, create and write the file 
func (api DefaultAPIService) WriteCSV() (string, error) {
	return api.repository.FindAll()
}

// NewAPIService to initialize the service
func NewAPIService(api domain.APIRepository) DefaultAPIService {
	return DefaultAPIService{repository: api}
}
