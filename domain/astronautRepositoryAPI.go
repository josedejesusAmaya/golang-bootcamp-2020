package domain

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// AstronautRepositoryAPI is
type AstronautRepositoryAPI struct {
	sb string
}

// FindAll is my function to create the CSV file
func (api AstronautRepositoryAPI) FindAll() (string, error) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		log.Fatalf("Error making HTTP request: %v", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error opening the file: %v", err)
	}

	api.sb = string(body)
	csvFile, err := os.Create("infrastructure/response.csv")
	if err != nil {
		log.Fatalf("Failed creating file: %v", err)
	}
	defer csvFile.Close()

	return api.sb, nil
}

// NewAstronautRepositoryAPI is
func NewAstronautRepositoryAPI() AstronautRepositoryAPI {
	return AstronautRepositoryAPI{sb: ""}
}
