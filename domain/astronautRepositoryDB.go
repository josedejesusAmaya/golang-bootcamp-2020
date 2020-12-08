package domain

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

// AstronautRepositoryDB is
type AstronautRepositoryDB struct {
	astronauts []Astronaut
}

// FindAll is my function to read the CSV file
func (a AstronautRepositoryDB) FindAll() ([]Astronaut, error) {
	file, err := os.Open("infrastructure/astronauts.csv")
	if err != nil {
		log.Fatalf("Error opening the file: %v", err)
		return nil, err
	}

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	defer file.Close()

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error reading line %v", err)
			return nil, err
		}

		hours, err := strconv.Atoi(record[13])
		if err != nil {
			log.Fatalf("Error %v", err)
			return nil, err
		}
		createAstronaut(&a.astronauts, record[0], hours)
	}
	return a.astronauts, nil
}

// NewAstronautRepositoryDB is the implementation of the helper
func NewAstronautRepositoryDB() AstronautRepositoryDB {
	astronauts := make([]Astronaut, 0)
	return AstronautRepositoryDB{astronauts: astronauts}
}

// Create my slice of astronauts
func createAstronaut(a *[]Astronaut, name string, hr int) {
	*a = append(*a, Astronaut{
		Name:     name,
		FlightHr: hr})
}
