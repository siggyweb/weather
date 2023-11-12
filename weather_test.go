package main

import (
	"reflect"
	"testing"
)

//TODO implement tests

func TestGetWeather(t *testing.T) {
	type args struct {
		l string
		w *Weather
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetWeather(tt.args.l, tt.args.w); (err != nil) != tt.wantErr {
				t.Errorf("GetWeather() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWeather_display(t *testing.T) {
	type fields struct {
		Location struct {
			Name           string
			Region         string
			Country        string
			Latitude       float64
			Longitude      float64
			Timezone       string
			LocaltimeEpock int
			LocalTime      string
		}
		Current struct {
			LastUpdated string
			TempC       float32
			Tempf       float32
			IsDay       int
			Condition   struct{ Description string }
		}
	}
	type args struct {
		flag bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := Weather{
				Location: tt.fields.Location,
				Current:  tt.fields.Current,
			}
			got, err := w.display(tt.args.flag)
			if (err != nil) != tt.wantErr {
				t.Errorf("Weather.display() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Weather.display() = %v, want %v", got, tt.want)
			}
		})
	}
}
