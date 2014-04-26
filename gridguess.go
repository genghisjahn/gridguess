package main

import "fmt"
import "flag"


var width = flag.Int("width", 10, "Horizontal width of the grid.")
var height = flag.Int("height", 10, "Vertical height of the grid.")


func main(){
	flag.Parse()
	fmt.Printf("This grid's width is %d.\n", *width)
	fmt.Printf("This grid's height is %d.\n", *height)
}