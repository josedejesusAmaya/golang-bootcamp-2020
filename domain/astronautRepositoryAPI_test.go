package domain

import (
	"os"
	"reflect"
	"testing"

	"github.com/josedejesusAmaya/golang-bootcamp-2020/errs"
)

func TestAstronautRepositoryAPI_FindAll(t *testing.T) {
	tests := []struct {
		name  string
		api   AstronautRepositoryAPI
		want  string
		want1 *errs.Error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.api.FindAll()
			if got != tt.want {
				t.Errorf("AstronautRepositoryAPI.FindAll() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("AstronautRepositoryAPI.FindAll() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNewAstronautRepositoryAPI(t *testing.T) {
	tests := []struct {
		name string
		want AstronautRepositoryAPI
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAstronautRepositoryAPI(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAstronautRepositoryAPI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_csvWriter(t *testing.T) {
	type args struct {
		file *os.File
		data []Response
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			csvWriter(tt.args.file, tt.args.data)
		})
	}
}
