package router

import (
	"encoding/csv"
	"io"
	"log"
	"net/http"
	"os"
)

const badRequest = 405

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
	f, err := os.Open("infrastructure/datastore/astronauts.csv")
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
		// fmt.Println(record[0]) Name
		// fmt.Println(record[13]) FlightHr
		data = append(data, record)
	}

}
