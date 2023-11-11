package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

// TODO json encoding tags.
type Weather struct {
	Location struct {
		Name           string  `json:"name"`
		Region         string  `json:"region"`
		Country        string  `json:"country"`
		Latitude       float64 `json:"lat"`
		Longitude      float64 `json:"lon"`
		Timezone       string  `json:"tz_id"`
		LocaltimeEpock int     `json:"localtime_epoch"`
		LocalTime      string  `json:"localtime"`
	}
	Current struct {
		LastUpdated string `json:"last_updated"`
		TempC       int    `json:"temp_c"`
		Tempf       int    `json:"temp_f"`
		IsDay       int    `json:"is_day"`
		Condition   struct {
			Description string `json:"text"`
			Icon        string `json:"icon"`
			Code        int    `json:"code"`
		}
	}
}

// Should the weather struct be broken out??

// TODO populate weather call stub.
func CallWeather(l string, w *Weather) error {
	// TODO retrieving secret could be modularised
	secret, err := os.ReadFile(".env")
	if err != nil {
		return fmt.Errorf("could not open secret file, %v", err)
	}

	// create url
	baseUrl := "https://api.weatherapi.com/v1/current"
	params := url.Values{}
	params.Add("key", string(secret))
	params.Add("location", l)
	url := baseUrl + params.Encode()

	// make the call to weather api and retrieve response
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("get call failed, %v", err)
	}
	rawJSON, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("could not read resonse body, %v", err)
	}

	// parse into a struct
	//weatherData := []byte{}
	json.Unmarshal(rawJSON, w)

	return nil
}
