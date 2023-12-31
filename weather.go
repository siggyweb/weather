package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
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
		LastUpdated string  `json:"last_updated"`
		TempC       float32 `json:"temp_c"`
		Tempf       float32 `json:"temp_f"`
		IsDay       int     `json:"is_day"`
		Condition   struct {
			Description string `json:"text"`
			// Icon        string `json:"icon"`
			// Code        int    `json:"code"`
		}
	}
}

// REfactor to recursively print struct.
func (w Weather) display(flag bool) ([]byte, error) {
	if flag {
		bytes, err := json.MarshalIndent(w, "", "\t")
		if err != nil {
			return nil, errors.New("could not marshall weather JSON")
		}
		return bytes, nil
	} else {
		str := fmt.Sprintf("The weather in %s is %s, with a temperature of %.01fC.\n", w.Location.Name, w.Current.Condition.Description, w.Current.TempC)
		return []byte(str), nil
	}
}

func GetWeather(l string, w *Weather) error {
	// TODO retrieving secret could be modularised
	secret, err := os.ReadFile(".env")
	if err != nil {
		return fmt.Errorf("could not open secret file, %v", err)
	}

	// TODO could modularise the request creation
	baseUrl := "https://api.weatherapi.com/v1/current.json"
	req, err := http.NewRequest(http.MethodGet, baseUrl, nil)
	if err != nil {
		return fmt.Errorf("request creation failed, %v", err)
	}

	req.Header.Set("accept", "application/json")
	q := req.URL.Query()
	q.Add("q", l)
	q.Add("key", string(secret))
	req.URL.RawQuery = q.Encode()

	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request failed, %v", err)
	}
	defer response.Body.Close()
	json.NewDecoder(response.Body).Decode(w)

	return nil
}
