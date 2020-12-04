package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

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
func HandleRequestAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Println("HandleRequestAPI")
	} else {
		w.WriteHeader(badRequest)
		w.Write([]byte("Method not allowed"))
	}
}

// HandleRequestRead to read the CSV file
func (a *AstronautHandler) HandleRequestRead(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		astronauts, _ := a.Service.GetAllAstronauts()
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(astronauts)
	} else {
		w.WriteHeader(badRequest)
		w.Write([]byte("Method not allowed"))
	}
}
