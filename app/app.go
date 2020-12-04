package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

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

// Start is the main function
func Start() {
	var cfg Config
	readConfig(&cfg)
	mux := http.NewServeMux()
	// wiring
	ah := handler.AstronautHandler{Service: service.NewAstronautService(domain.NewAstronautRepositoryStub())}
	mux.HandleFunc("/api/astronauts", ah.HandleRequestAPI)
	mux.HandleFunc("/api/read", handler.HandleRequestRead)
	fmt.Println("Service is running")
	err := http.ListenAndServe(cfg.Server.Port, mux)
	if err != nil {
		log.Fatalf("Failed to listen on port 8000: %v", err)
		return
	}
}

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
