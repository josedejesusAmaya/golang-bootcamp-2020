package service

import (
	"reflect"
	"testing"

	"github.com/josedejesusAmaya/golang-bootcamp-2020/domain"
	"github.com/josedejesusAmaya/golang-bootcamp-2020/errs"
)

func TestDefaultAstronautService_GetAscAstronauts(t *testing.T) {
	tests := []struct {
		name  string
		d     DefaultAstronautService
		want  []domain.Astronaut
		want1 *errs.Error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.d.GetAscAstronauts()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultAstronautService.GetAscAstronauts() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("DefaultAstronautService.GetAscAstronauts() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDefaultAstronautService_GetDescAstronauts(t *testing.T) {
	tests := []struct {
		name  string
		d     DefaultAstronautService
		want  []domain.Astronaut
		want1 *errs.Error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.d.GetDescAstronauts()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultAstronautService.GetDescAstronauts() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("DefaultAstronautService.GetDescAstronauts() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNewAstronautService(t *testing.T) {
	type args struct {
		a domain.AstronautRepository
	}
	tests := []struct {
		name string
		args args
		want DefaultAstronautService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAstronautService(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAstronautService() = %v, want %v", got, tt.want)
			}
		})
	}
}
