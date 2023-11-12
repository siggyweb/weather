package main

import (
	"fmt"
	"os"
)

func main() {
	// TODO input guarding for multiple inputs
	input := os.Args
	if len(input) <= 0 {
		fmt.Println("please enter a location string")
		os.Exit(0)
	}

	// TODO check for verbose flag

	w := Weather{}
	GetWeather(input[1], &w)
	
	pretty, err := w.display()
	if err == nil {
		fmt.Println(string(pretty))
	} else {
		fmt.Println(err, " exiting.")
		os.Exit(1)
	}
	
}
