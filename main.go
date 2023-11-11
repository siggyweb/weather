package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("The weather is sunny.")

	// TODO take command line arg for weather and collect data.
	// change this back to 1 for released version
	input := os.Args
	if len(input) <= 0 {
		fmt.Println("please enter a location string")
		os.Exit(0)
	}
	w := Weather{}
	CallWeather("London", &w)
	fmt.Println(w)
}
