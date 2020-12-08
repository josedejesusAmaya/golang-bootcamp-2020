package service

import "github.com/josedejesusAmaya/golang-bootcamp-2020/domain"

// APIService is
type APIService interface {
	WriteCSV() (string, error)
}

// DefaultAPIService is
type DefaultAPIService struct {
	repository domain.APIRepository
}

// WriteCSV is
func (api DefaultAPIService) WriteCSV() (string, error) {
	return api.repository.FindAll()
}

// NewAPIService is
func NewAPIService(api domain.APIRepository) DefaultAPIService {
	return DefaultAPIService{repository: api}
}
