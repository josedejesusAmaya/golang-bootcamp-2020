package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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
	fmt.Println("HandleRequestAPI")
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
	fmt.Println("HandleRequestDB")
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

// HandleRequest to handle the HTTP request
func HandleRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("Request received:", r.Method)
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
