package domain

import (
	"net/http"
	"testing"
)

func TestAstronautRepositoryDB_FindAll_error_opening_file(t *testing.T) {
	var db AstronautRepositoryDB

	_, apiError := db.FindAll()

	if apiError.Message != "Error opening the file" {
		t.Error("Invalid message FindAll error while opening file")
	}

	if apiError.Code != http.StatusInternalServerError {
		t.Error("Invalid code FindAll error while opening file")
	}
}

func Test_NewAstronautRepositoryDB_validating_return(t *testing.T) {
	var list AstronautRepositoryDB

	resp := NewAstronautRepositoryDB()

	if len(resp.astronauts) != len(list.astronauts) {
		t.Error("Invalid return from NewAstronautRepositoryDB")
	}
}

func Test_FindAsc_validating_return(t *testing.T) {
	var ar AstronautRepositoryDB
	_, apiError := ar.FindAsc()

	if apiError != nil {
		t.Error("Invalid return from FindAsc")
	}
}

func Test_FindDesc_validating_return(t *testing.T) {
	var ar AstronautRepositoryDB
	_, apiError := ar.FindDesc()

	if apiError != nil {
		t.Error("Invalid return from FindDesc")
	}
}
