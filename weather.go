package main

// TODO json encoding tags.
type Weather struct {
	Location struct {
		Name           string
		Region         string
		Country        string
		Lat            float64
		Long           float64
		Tz_id          string
		LocaltimeEpock int
		Localtime      string
	}
	Current struct {
		LastUpdated string
		TempC       int
		Tempf       int
		IsDay       int
		Condition   struct {
			Desc string
			Icon string
			Code int
		}
	}
}
