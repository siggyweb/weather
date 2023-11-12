package main

import (
	"fmt"
	"os"
)

func main() {
	// TODO input guarding for multiple inputs
	input := os.Args
	if len(input) <= 1 {
		fmt.Println("please enter a location string in the form \"weather <location> flag[optional]\". See README for more info. ")
		os.Exit(1)
	}

	w := Weather{}
	GetWeather(input[1], &w)

	var verboseFlag bool
	if len(input) == 3 && input[2] == "-v" {
		verboseFlag = true
	}

	pretty, err := w.display(verboseFlag)
	if err == nil {
		fmt.Println(string(pretty))
	} else {
		fmt.Println(err, " exiting.")
		os.Exit(1)
	}

}
