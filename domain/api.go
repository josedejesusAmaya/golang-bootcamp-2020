package domain

// API is the main structure of the response from creating the CSV file
type API struct {
	Message int `json:"message"`
}

// APIRepository is is a port from the API
type APIRepository interface {
	FindAll() (string, error)
}
