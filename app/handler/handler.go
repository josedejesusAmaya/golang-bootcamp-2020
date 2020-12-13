package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
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

// Wiring is
type Wiring struct {
	DB  AstronautHandler
	API APIHandler
}

// requestAPI to consume external API and write the CSV
func requestAPI(w http.ResponseWriter, r *http.Request, myWiring Wiring) {
	message, err := myWiring.API.Service.WriteCSV()
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, message)
	}
}

// requestDB to read the CSV file
func requestDB(w http.ResponseWriter, r *http.Request, myWiring Wiring) {
	astronauts, err := myWiring.DB.Service.GetAllAstronauts()
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, astronauts)
	}
}

// HandleRequest to handle the HTTP request to GET and read the astronauts
func (wiring Wiring) HandleRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		requestAPI(w, r, wiring)
	case "GET":
		requestDB(w, r, wiring)
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Method not allowed"))
	}
}

// OrderedHandleRequest to get an ordered list of astronauts
func (wiring Wiring) OrderedHandleRequest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["order"] == "asc" {
		ascAstronautsList(w, r, wiring)
	}

	if vars["order"] == "desc" {
		descAstronautsList(w, r, wiring)
	}
}

// ascAstronautsList returns asc list by the spaceFlightHr field
func ascAstronautsList(w http.ResponseWriter, r *http.Request, myWiring Wiring) {
	astronauts, err := myWiring.DB.Service.GetAscAstronauts()
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, astronauts)
	}
}

// descAstronautsList returns desc list by the spaceFlightHr field
func descAstronautsList(w http.ResponseWriter, r *http.Request, myWiring Wiring) {
	astronauts, err := myWiring.DB.Service.GetDescAstronauts()
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, astronauts)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
