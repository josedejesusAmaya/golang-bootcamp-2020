package domain

// Astronaut is the main structure for collecting data from the CSV
type Astronaut struct {
	Name     string `json:"name"`
	FlightHr int    `json:"flightHr"`
}

// AstronautRepository is a port from the DB
type AstronautRepository interface {
	FindAll() ([]Astronaut, error)
}
