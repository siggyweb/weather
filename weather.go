package main

// TODO json encoding tags.
type Weather struct {
	Location struct {
		Name           string  `json:"name"`
		Region         string  `json:"region"`
		Country        string  `json:"country"`
		Latitude            float64 `json:"lat"`
		Longitude           float64 `json:"lon"`
		Timezone          string  `json:"tz_id"`
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
			Icon string `json:"icon"`
			Code int    `json:"code"`
		}
	}
}
