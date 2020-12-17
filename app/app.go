package app

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/josedejesusAmaya/golang-bootcamp-2020/app/handler"
	"github.com/josedejesusAmaya/golang-bootcamp-2020/domain"
	"github.com/josedejesusAmaya/golang-bootcamp-2020/service"
	"gopkg.in/yaml.v2"
)

// Config is a type of config.yml file
type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`
}

var wiring handler.Wiring

// Start is the main function of the app
func Start() {
	var cfg Config
	router := mux.NewRouter()
	readConfig(&cfg)
	wiring.DB = handler.AstronautHandler{Service: service.NewAstronautService(domain.NewAstronautRepositoryDB())}
	wiring.API = handler.APIHandler{Service: service.NewAPIService(domain.NewAstronautRepositoryAPI())}
	router.HandleFunc("/api/astronauts", wiring.HandleRequest)
	router.HandleFunc("/api/astronauts/{order:[a-z]+}", wiring.OrderedHandleRequest)
	log.Print("Service is running")
	err := http.ListenAndServe(cfg.Server.Port, router)
	if err != nil {
		log.Fatalf("Failed to listen on port 8000: %v", err)
		return
	}
}

// readConfig to decode yml file
func readConfig(cfg *Config) {
	f, err := os.Open("config.yml")
	if err != nil {
		log.Fatalf("Failed to read yml file: %v", err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		log.Fatalf("Failed to decode yml: %v", err)
	}
}
