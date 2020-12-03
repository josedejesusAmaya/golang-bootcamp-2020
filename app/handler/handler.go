package handler

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/josedejesusAmaya/golang-bootcamp-2020/app/domain/model"
)

const badRequest = 405

// Astronaut object to read the CSV
type Astronaut = model.Astronaut

// HandleRequestRead is the handler of my routes
func HandleRequestRead(w http.ResponseWriter, r *http.Request) {
	log.Println("Request Read received:", r.Method)
	if r.Method == "GET" {
		readCSV(w, r)
	} else {
		w.WriteHeader(badRequest)
		w.Write([]byte("Method not allowed"))
	}
}

func readCSV(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("app/infrastructure/datastore/astronauts.csv")
	if err != nil {
		log.Fatalf("Error opening the file: %v", err)
	}

	defer f.Close()
	reader := csv.NewReader(f)
	reader.FieldsPerRecord = -1

	var data [][]string

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error reading line %v", err)
		}

		createAstronaut(record[0], record[13])
		data = append(data, record)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func createAstronaut(name string, hr string) {
	// fmt.Println(name, hr)
}
