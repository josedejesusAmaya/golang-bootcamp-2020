package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/gorilla/mux"
	"github.com/josedejesusAmaya/golang-bootcamp-2020/domain"
	"github.com/josedejesusAmaya/golang-bootcamp-2020/errs"
	"github.com/josedejesusAmaya/golang-bootcamp-2020/mocks/service"
)

func Test_requestAPI_with_status_code_200(t *testing.T) {
	var wiring Wiring
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := service.NewMockAPIService(ctrl)
	mockService.EXPECT().WriteCSV().Return("CSV file was created", nil)
	wiring.API = APIHandler{Service: mockService}
	router := mux.NewRouter()
	router.HandleFunc("/api/astronauts", wiring.HandleRequest)
	request, _ := http.NewRequest(http.MethodPost, "/api/astronauts", nil)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Error("Failed while testing requestDB with status code 200")
	}
}

func Test_requestAPI_with_status_code_500_with_error_message(t *testing.T) {
	var wiring Wiring
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := service.NewMockAPIService(ctrl)
	mockService.EXPECT().WriteCSV().Return("", errs.NewUnexpectedError("Unexpected database error"))
	wiring.API = APIHandler{Service: mockService}
	router := mux.NewRouter()
	router.HandleFunc("/api/astronauts", wiring.HandleRequest)
	request, _ := http.NewRequest(http.MethodPost, "/api/astronauts", nil)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusInternalServerError {
		t.Error("Failed while testing requestDB with status code 500")
	}
}

func Test_requestDB_with_status_code_200(t *testing.T) {
	var wiring Wiring
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := service.NewMockAstronautService(ctrl)
	dummyAstronauts := []domain.Astronaut{
		{Name: "Jane Doe", FlightHr: 30},
		{Name: "Pepe Amaya", FlightHr: 19},
		{Name: "John Doe", FlightHr: 10},
	}
	mockService.EXPECT().GetAllAstronauts().Return(dummyAstronauts, nil)
	wiring.DB = AstronautHandler{Service: mockService}
	router := mux.NewRouter()
	router.HandleFunc("/api/astronauts", wiring.HandleRequest)
	request, _ := http.NewRequest(http.MethodGet, "/api/astronauts", nil)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Error("Failed while testing requestDB with status code 200")
	}
}

func Test_requestDB_with_status_code_500_with_error_message(t *testing.T) {
	var wiring Wiring
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := service.NewMockAstronautService(ctrl)
	mockService.EXPECT().GetAllAstronauts().Return(nil, errs.NewUnexpectedError("Unexpected database error"))
	wiring.DB = AstronautHandler{Service: mockService}
	router := mux.NewRouter()
	router.HandleFunc("/api/astronauts", wiring.HandleRequest)
	request, _ := http.NewRequest(http.MethodGet, "/api/astronauts", nil)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusInternalServerError {
		t.Error("Failed while testing requestDB with status code 500")
	}
}
