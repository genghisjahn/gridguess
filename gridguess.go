package main

import "fmt"
import "flag"
import "math/rand"
import "time"

var width = flag.Int("width", 10, "Horizontal width of the grid. Default is 10.")
var height = flag.Int("height", 10, "Vertical height of the grid. Default is 10.")

func main() {
	flag.Parse()

	rand.Seed(time.Now().UTC().UnixNano())

	grid := Grid{}
	grid.Build()

	fmt.Printf("Enter Guess: ")
	var guess_str string
	_, err := fmt.Scanf("%v", &guess_str)
	if err != nil {
		fmt.Printf("Error: %v.\n", err)
	}
	gresult, err_result := grid.ProcessGuess(guess_str)
	if err_result != nil {
		fmt.Printf("Error: %v.\n", err_result)
	}
	fmt.Printf("%v.\n",gresult)

}
