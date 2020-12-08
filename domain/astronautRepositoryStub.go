package domain

// AstronautRepositoryStub is a truct for the data testing
type AstronautRepositoryStub struct {
	astronauts []Astronaut
}

// FindAll is a function for testing
func (a AstronautRepositoryStub) FindAll() ([]Astronaut, error) {
	return a.astronauts, nil
}

// NewAstronautRepositoryStub is a function for double testing
func NewAstronautRepositoryStub() AstronautRepositoryStub {
	astronauts := []Astronaut{
		{Name: "Jane Doe", FlightHr: 30},
		{Name: "Pepe Amaya", FlightHr: 19},
		{Name: "John Doe", FlightHr: 10},
	}
	return AstronautRepositoryStub{astronauts: astronauts}
}
