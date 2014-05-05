package main

import "fmt"
import "math/rand"
import "time"
import "github.com/daviddengcn/go-colortext"
import "flag"

var length = flag.Int("length", 10, "Length for each dimension.  Default is 10.")

func main() {
	flag.Parse()
	rand.Seed(time.Now().UTC().UnixNano())
	grid := Grid{}
	grid.Build(*length)
	PlayLoop(grid)
}
func PlayLoop(grid Grid) {
	found := false
	grid.DescribeSpace()
	fmt.Printf("\n")
	for !found {
		ct.ChangeColor(ct.Blue, true, ct.Black, false)
		fmt.Printf("\nEnter Guess: ")
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
		if gresult.Found == true {
			found = true
			ct.ChangeColor(ct.Green, true, ct.Black, false)
		}
		fmt.Printf("%v\n\n", gresult)
		ct.ChangeColor(ct.Green, true, ct.Black, false)
		grid.DescribeSpace()
	}
	fmt.Printf("\n\n")
}
