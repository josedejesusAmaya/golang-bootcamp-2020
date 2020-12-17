package domain

import "github.com/josedejesusAmaya/golang-bootcamp-2020/errs"

// AstronautRepositoryStub is the type of the state for the data testing
type AstronautRepositoryStub struct {
	Astronauts []Astronaut
}

// FindAll is a function for testing
func (a AstronautRepositoryStub) FindAll() ([]Astronaut, *errs.Error) {
	return a.Astronauts, nil
}

// NewAstronautRepositoryStub is a function for double testing
func NewAstronautRepositoryStub() AstronautRepositoryStub {
	astronauts := []Astronaut{
		{Name: "Jane Doe", FlightHr: 30},
		{Name: "Pepe Amaya", FlightHr: 19},
		{Name: "John Doe", FlightHr: 10},
	}
	return AstronautRepositoryStub{Astronauts: astronauts}
}
