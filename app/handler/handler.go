package handler

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/josedejesusAmaya/golang-bootcamp-2020/domain"
	"github.com/josedejesusAmaya/golang-bootcamp-2020/service"
)

const badRequest = 405

// Astronaut object to read the CSV
type Astronaut = domain.Astronaut

// AstronautHandler is
type AstronautHandler struct {
	Service service.AstronautService
}

// HandleRequestAPI to consume external API and write the CSV
func (a *AstronautHandler) HandleRequestAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		astronauts, _ := a.Service.GetAllAstronauts()
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(astronauts)
		// writeCSV(w, r)
	} else {
		w.WriteHeader(badRequest)
		w.Write([]byte("Method not allowed"))
	}
}

// HandleRequestRead to read my CSV
func HandleRequestRead(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		readCSV(w, r)
	} else {
		w.WriteHeader(badRequest)
		w.Write([]byte("Method not allowed"))
	}
}

func writeCSV(w http.ResponseWriter, r *http.Request) {

}

func readCSV(w http.ResponseWriter, r *http.Request) {
	astronauts := make([]Astronaut, 0)
	f, err := os.Open("infrastructure/astronauts.csv")
	if err != nil {
		log.Fatalf("Error opening the file: %v", err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.FieldsPerRecord = -1
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error reading line %v", err)
		}

		hours, err := strconv.Atoi(record[13])
		if err != nil {
			log.Fatalf("Error %v", err)
		}
		createAstronaut(&astronauts, record[0], hours)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(astronauts)
}

func createAstronaut(a *[]Astronaut, name string, hr int) {
	*a = append(*a, Astronaut{
		Name:     name,
		FlightHr: hr})
}
