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

// TODO create a weather receiver to display the info to the console

func CallWeather(l string, w *Weather) error {
	// TODO retrieving secret could be modularised
	secret, err := os.ReadFile(".env")
	if err != nil {
		return fmt.Errorf("could not open secret file, %v", err)
	}

	//TODO refactor verbose call

	// way 1 - verbose
	baseUrl := "https://api.weatherapi.com/v1/current.json?"
	params := url.Values{}
	params.Add("q", l)
	params.Add("key", string(secret))
	url := baseUrl + params.Encode()

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("get call failed, %v", err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("get call failed, %v", err)
	}
	fmt.Println(string(body))

	// way 2 - brief
	baseUrl2 := "https://api.weatherapi.com/v1/current.json"
	req, err := http.NewRequest(http.MethodGet, baseUrl2, nil)
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
