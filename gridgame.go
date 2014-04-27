package main

import "fmt"
import "math/rand"
import "time"
import "github.com/daviddengcn/go-colortext"

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
