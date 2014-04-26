package main

import "fmt"
import "flag"


var width = flag.Int("width", 10, "How wide the grid is.")


func main(){
	flag.Parse()
	fmt.Printf("This grid width is %d.\n", *width)
}