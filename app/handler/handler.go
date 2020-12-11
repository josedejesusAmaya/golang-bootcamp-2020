package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/josedejesusAmaya/golang-bootcamp-2020/domain"
	"github.com/josedejesusAmaya/golang-bootcamp-2020/service"
)

// AstronautHandler is a type to create astronaut services
type AstronautHandler struct {
	Service service.AstronautService
}

// APIHandler is a type to create api services
type APIHandler struct {
	Service service.APIService
}

// requestAPI to consume external API and write the CSV
func requestAPI(w http.ResponseWriter, r *http.Request) {
	// wiring
	api := APIHandler{Service: service.NewAPIService(domain.NewAstronautRepositoryAPI())}
	message, err := api.Service.WriteCSV()
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(message)
	}
}

// requestDB to read the CSV file
func requestDB(w http.ResponseWriter, r *http.Request) {
	// wiring
	db := AstronautHandler{Service: service.NewAstronautService(domain.NewAstronautRepositoryDB())}
	astronauts, err := db.Service.GetAllAstronauts()
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(astronauts)
	}
}

// HandleRequest to handle the HTTP request to GET and read the astronauts
func HandleRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		requestAPI(w, r)
	case "GET":
		requestDB(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Method not allowed"))
	}
}

// OrderedHandleRequest to get an ordered list of astronauts
func OrderedHandleRequest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["order"] == "asc" {
		ascAstronautsList()
	}

	if vars["order"] == "desc" {
		descAstronautsList()
	}
}

// ascAstronautsList returns asc list by the spaceFlightHr field
func ascAstronautsList() {

}

// descAstronautsList returns desc list by the spaceFlightHr field
func descAstronautsList() {

}
