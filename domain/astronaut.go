package domain

import "github.com/josedejesusAmaya/golang-bootcamp-2020/errs"

// Astronaut is the main structure for collecting data from the CSV
type Astronaut struct {
	Name     string `json:"name"`
	FlightHr int    `json:"flightHr"`
}

// AstronautRepository is a port from the DB
type AstronautRepository interface {
	FindAll() ([]Astronaut, *errs.Error)
	FindAsc() ([]Astronaut, *errs.Error)
	FindDesc() ([]Astronaut, *errs.Error)
}
