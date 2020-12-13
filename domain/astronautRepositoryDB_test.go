package domain

import (
	"reflect"
	"testing"
)

func TestAstronautRepositoryDB_FindAll(t *testing.T) {
	tests := []struct {
		name    string
		a       AstronautRepositoryDB
		want    []Astronaut
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.FindAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("AstronautRepositoryDB.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AstronautRepositoryDB.FindAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAstronautRepositoryDB(t *testing.T) {
	tests := []struct {
		name string
		want AstronautRepositoryDB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAstronautRepositoryDB(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAstronautRepositoryDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createAstronaut(t *testing.T) {
	type args struct {
		a    *[]Astronaut
		name string
		hr   int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createAstronaut(tt.args.a, tt.args.name, tt.args.hr)
		})
	}
}
