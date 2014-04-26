package main

import "fmt"
import "flag"


var width = flag.Int("width", 10, "Horizontal width of the grid.")
var height = flag.Int("height", 10, "Vertical height of the grid.")


func main(){
	flag.Parse()

	point:=Point{}
	point.x=*width
	point.y=*height
	fmt.Printf("Here is the point %v.\n",point)
}