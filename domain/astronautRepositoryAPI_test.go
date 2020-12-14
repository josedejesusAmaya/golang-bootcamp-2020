package domain

import (
	"net/http"
	"testing"
)

func TestAstronautRepositoryAPI_FindAll_error_creating_file(t *testing.T) {
	api := NewAstronautRepositoryAPI()

	_, apiError := api.FindAll()

	if apiError.Message != "Failed to create file" {
		t.Error("Invalid message FindAll error while creating file")
	}

	if apiError.Code != http.StatusInternalServerError {
		t.Error("Invalid code FindAll error while creating file")
	}
}
