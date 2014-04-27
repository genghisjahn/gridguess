package main

import "fmt"
import "flag"
import "math/rand"
import "time"
import "github.com/daviddengcn/go-colortext"

var width = flag.Int("width", 10, "Horizontal width of the grid. Default is 10.")
var height = flag.Int("height", 10, "Vertical height of the grid. Default is 10.")
var length = flag.Int("length", -1, "Length/Depth of the grid.  Default is -1.")


func main() {
	flag.Parse()

	rand.Seed(time.Now().UTC().UnixNano())

	grid := Grid{}
	grid.Build()

	PlayLoop(grid)

}
func PlayLoop(grid Grid) {
	found := false
	for !found {
		ct.ChangeColor(ct.Blue, true, ct.Black, false)
		fmt.Printf("\nEnter Guess:")
		var guess_str string
		_, err := fmt.Scanf("%v", &guess_str)
		if err != nil {
			ct.ChangeColor(ct.Red, true, ct.Black, false)
			fmt.Printf("Error: %v.\n", err)
		}
		gresult, err_result := grid.ProcessGuess(guess_str)
		if err_result != nil {
			ct.ChangeColor(ct.Red, true, ct.Black, false)
			fmt.Printf("%v.\n", err_result)
		}
		ct.ChangeColor(ct.Yellow, true, ct.Black, false)
		if gresult.VerticalPosition == cFound && gresult.HorizontalPosition == cFound {
			found = true
			ct.ChangeColor(ct.Green, true, ct.Black, false)
		}
		fmt.Printf("%v\n", gresult)
	}
	fmt.Printf("\n\n")
}
