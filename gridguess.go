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

	fmt.Printf("Num of Points %v.\n", len(grid.Points))

}
