# Golang Bootcamp

The API brings information about astronauts from the NASA. Storage the response to a CSV file and then read that information, sort it, and present it as JSON. In addition, I included two more requests to sort this list ascending and descending by the flight hours.

* API: https://astronautsapinodejs.herokuapp.com/astronauts
* Data extracted from: https://www.kaggle.com/nasa/astronaut-yearbook?select=astronauts.csv

## Installation
Clone the project via ssh or https
```bash
git clone https://github.com/josedejesusAmaya/golang-bootcamp-2020.git
```

### How to run
```bash
go run main.go
```

Send requests through the routes.rest file.

## Usage for Test
```bash
go test ./...
```

## About
* Language: GO
* Extensions: REST Client
* Testing: GoMock
