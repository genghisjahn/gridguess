package main

import "fmt"
import "flag"

var width = flag.Int("width", 10, "Horizontal width of the grid.")
var height = flag.Int("height", 10, "Vertical height of the grid.")

func main() {
	flag.Parse()

	point := Point{}
	point.x = *width
	point.y = *height
	fmt.Printf("Here is the point %v.\n", point)

	grid := Grid{}
	grid.Points = make([]Point, 100)

	v := 1
	h := 1
	for v <= *width {
		for h <= *height {
			tempPoint := Point{}
			tempPoint.x = h
			tempPoint.y = v
			h += 1
			fmt.Printf("Here is the point %v.\n", tempPoint)
		}
		v += 1
		h = 1
	}

}
