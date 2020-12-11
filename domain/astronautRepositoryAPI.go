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

// AstronautRepositoryAPI is the type of the state
type AstronautRepositoryAPI struct {
	sb string
}

// FindAll is my function to create the CSV file
func (api AstronautRepositoryAPI) FindAll() (string, error) {
	resp, err := http.Get("https://astronautsapinodejs.herokuapp.com/astronauts")
	if err != nil {
		log.Fatalf("Error making HTTP request: %v", err)
		return "", err
	}
	jsonDataFromFile, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error opening the file: %v", err)
		return "", err
	}
	var jsonData []Response
	err = json.Unmarshal([]byte(jsonDataFromFile), &jsonData)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	csvFile, err := os.Create("infrastructure/astronauts.csv")
	if err != nil {
		log.Fatalf("Failed creating file: %v", err)
		return "Error ", err
	}
	defer csvFile.Close()
	csvWriter(csvFile, jsonData)
	return "message: CSV file was created", nil
}

// NewAstronautRepositoryAPI to initialize repository
func NewAstronautRepositoryAPI() AstronautRepositoryAPI {
	return AstronautRepositoryAPI{sb: ""}
}

// csvWriter to write information of the HTTP response
func csvWriter(file *os.File, data []Response) {
	writer := csv.NewWriter(file)
	for _, usance := range data {
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
}
