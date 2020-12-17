# Golang Bootcamp

The API handles information about American astronauts from 1959-2016. I use an external API of a service that I create with NodeJS on Heroku.This project storage the response into a CSV file and then read that information, sort it, and present it as JSON. In addition, I included two more requests to sort this list ascending and descending by the flight hours.
Use the routes.rest file to test the endpoints.

* External API: https://astronautsapinodejs.herokuapp.com/astronauts
* Data extracted from: https://www.kaggle.com/nasa/astronaut-yearbook?select=astronauts.csv

## Installation
Clone the project via ssh or https
```bash
git clone https://github.com/josedejesusAmaya/golang-bootcamp-2020.git
```

```bash
cd golang-bootcamp-2020/
go build
```

### How to run
```bash
go run main.go
```

The project runs a web server on `port 8080`

## Usage for Test
```bash
go test ./...
```

## About
* Language: GO
* Extensions: REST Client
* Testing: GoMock
