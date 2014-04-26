package main

import "fmt"
import "flag"
import "math/rand"
import "time"

var width = flag.Int("width", 10, "Horizontal width of the grid. Default is 10.")
var height = flag.Int("height", 10, "Vertical height of the grid. Default is 10.")



func main() {
	flag.Parse()

	rand.Seed( time.Now().UTC().UnixNano())

	grid := Grid{}

	target_x := randInt(1,*width)
	target_y := randInt(1,*height)

	set_x := 1
	set_y := 1
	for set_x <= *width {
		for set_y <= *height {
			tempPoint := Point{}
			tempPoint.x = set_x
			tempPoint.y = set_y
			set_y += 1

			if tempPoint.x == target_x && tempPoint.y == target_y {
				tempPoint.target=true
			}

			grid.Points = append(grid.Points, tempPoint)
			fmt.Printf("Here is the point %v.\n", tempPoint)
		}
		set_x += 1
		set_y = 1
	}
	fmt.Printf("Num of Points %v.\n", len(grid.Points))
	fmt.Printf("Target (X,Y) (%d,%d)",target_x,target_y)
}
