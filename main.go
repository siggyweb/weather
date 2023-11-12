package main

import (
	"fmt"
	"os"
)

func main() {
	// TODO input guarding for multiple inputs
	input := os.Args
	if len(input) <= 1 {
		fmt.Println("please enter a location string")
		os.Exit(0)
	}

	// TODO check for verbose flag

	w := Weather{}
	CallWeather(input[1], &w)
	fmt.Println(w)
}
