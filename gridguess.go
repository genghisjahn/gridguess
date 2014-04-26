package main

import "fmt"
import "flag"


var width = flag.Int("width", 12, "How wide the grid is.")


func main(){

	// fmt.Printf("You're here.\n")
	fmt.Printf("This grid width is %d.\n", *width)

	num:=12
	fmt.Printf("The number is %d.\n",num)
}