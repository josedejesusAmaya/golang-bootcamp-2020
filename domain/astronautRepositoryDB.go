package domain

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"

	"github.com/josedejesusAmaya/golang-bootcamp-2020/errs"
)

// AstronautRepositoryDB is the type of the state
type AstronautRepositoryDB struct {
	astronauts []Astronaut
}

var originalList []Astronaut

// FindAll is my function to read the CSV file and return all astronauts
func (a AstronautRepositoryDB) FindAll() ([]Astronaut, *errs.Error) {
	file, err := os.Open("infrastructure/astronauts.csv")
	if err != nil {
		return nil, errs.NewUnexpectedError("Error opening the file")
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
			return nil, errs.NewUnexpectedError("Error reading line")
		}

		hours, err := strconv.Atoi(record[13])
		if err != nil {
			return nil, errs.NewUnexpectedError("Error casting int to string")
		}
		createAstronaut(&originalList, record[0], hours)
	}
	a.astronauts = originalList
	return a.astronauts, nil
}

// FindAsc to return all asc ordered astronauts
func (a AstronautRepositoryDB) FindAsc() ([]Astronaut, *errs.Error) {
	auxList := originalList
	count := 0
	change := true
	for change {
		change = false
		for i := 1; i < (len(auxList) - count); i++ {
			if auxList[i-1].FlightHr > auxList[i].FlightHr {
				change = true
				orderList(&auxList, i)
			}
		}
		count++
	}
	a.astronauts = auxList
	return a.astronauts, nil
}

// orderList swap the FlighHr values to sort the astronauts list
func orderList(a *[]Astronaut, right int) {
	left := right - 1
	newList := *a
	aux := newList[left].FlightHr
	newList[left].FlightHr = newList[right].FlightHr
	newList[right].FlightHr = aux
}

// FindDesc is to return all desc ordered astronauts
func (a AstronautRepositoryDB) FindDesc() ([]Astronaut, *errs.Error) {
	auxList := originalList
	count := 0
	change := true
	for change {
		change = false
		for i := 1; i < (len(auxList) - count); i++ {
			if auxList[i-1].FlightHr < auxList[i].FlightHr {
				change = true
				orderList(&auxList, i)
			}
		}
		count++
	}
	a.astronauts = auxList
	return a.astronauts, nil
}

// NewAstronautRepositoryDB is the implementation of the helper
func NewAstronautRepositoryDB() AstronautRepositoryDB {
	originalList := make([]Astronaut, 0)
	return AstronautRepositoryDB{astronauts: originalList}
}

// Create my slice of astronauts
func createAstronaut(a *[]Astronaut, name string, hr int) {
	*a = append(*a, Astronaut{
		Name:     name,
		FlightHr: hr})
}
