package domain

// API is the main structure of the response from creating the CSV file
type API struct {
	Message int `json:"message"`
}

// APIRepository is is a port from the API
type APIRepository interface {
	FindAll() (string, error)
}

// Response is the type  
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
