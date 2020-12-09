package domain

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

// AstronautRepositoryAPI is
type AstronautRepositoryAPI struct {
	sb string
}

// Response is
type Response struct {
	Name               string  `json:"name"`
	Year               int     `json:"year"`
	Group              int     `json:"group"`
	Status             string  `json:"status"`
	BirthDate          string  `json:"birthDate"`
	BirthPlace         string  `json:"birthPlace"`
	Gender             string  `json:"gender"`
	AlmaMater          string  `json:"almaMater"`
	UndergraduateMajor string  `json:"undergraduateMajor"`
	GraduateMajor      string  `json:"graduateMajor"`
	MilitaryRank       string  `json:"militaryRank"`
	MilitaryBranch     string  `json:"militaryBranch"`
	SpaceFlights       int     `json:"spaceFlights"`
	SpaceFlightHr      float64 `json:"spaceFlightHr"`
	SpaceWalks         int     `json:"spaceWalks"`
	SpaceWalksHr       float64 `json:"spaceWalksHr"`
	Missions           string  `json:"missions"`
	DeathDate          string  `json:"deathDate"`
	DeathMission       string  `json:"deathMission"`
}

// FindAll is my function to create the CSV file
func (api AstronautRepositoryAPI) FindAll() (string, error) {
	resp, err := http.Get("http://localhost:3000/api/astronauts")
	if err != nil {
		log.Fatalf("Error making HTTP request: %v", err)
		return "", err
	}
	jsonDataFromFile, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error opening the file: %v", err)
		return "", err
	}
	/*jsonDataFromFile, err := ioutil.ReadFile("infrastructure/astronauts.json")
	if err != nil {
		log.Fatalf("Error while reading JSON: %v", err)
		return "", err
	}*/
	var jsonData []Response
	err = json.Unmarshal([]byte(jsonDataFromFile), &jsonData)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	// api.sb = string(body)
	csvFile, err := os.Create("infrastructure/astronauts.csv")
	if err != nil {
		log.Fatalf("Failed creating file: %v", err)
		return "Error ", err
	}
	defer csvFile.Close()
	writer := csv.NewWriter(csvFile)
	for _, usance := range jsonData {
		var row []string
		row = append(row, usance.Name)
		year := strconv.Itoa(usance.Year)
		row = append(row, year)
		group := strconv.Itoa(usance.Group)
		row = append(row, group)
		row = append(row, usance.Status)
		row = append(row, usance.BirthDate)
		row = append(row, usance.BirthPlace)
		row = append(row, usance.Gender)
		row = append(row, usance.AlmaMater)
		row = append(row, usance.UndergraduateMajor)
		row = append(row, usance.GraduateMajor)
		row = append(row, usance.MilitaryRank)
		row = append(row, usance.MilitaryBranch)
		spaceFlights := strconv.Itoa(usance.SpaceFlights)
		row = append(row, spaceFlights)
		spaceFlightHr := strconv.FormatFloat(usance.SpaceFlightHr, 'f', -1, 64)
		row = append(row, spaceFlightHr)
		spaceWalks := strconv.Itoa(usance.SpaceWalks)
		row = append(row, spaceWalks)
		spaceWalksHr := strconv.FormatFloat(usance.SpaceWalksHr, 'f', -1, 64)
		row = append(row, spaceWalksHr)
		row = append(row, usance.Missions)
		row = append(row, usance.DeathDate)
		row = append(row, usance.DeathMission)
		writer.Write(row)
	}
	writer.Flush()
	return "message: CSV file was created", nil
}

// NewAstronautRepositoryAPI is
func NewAstronautRepositoryAPI() AstronautRepositoryAPI {
	return AstronautRepositoryAPI{sb: ""}
}
