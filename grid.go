package main

import "math/rand"
import "strings"
import "errors"
//import "strconv"
import "fmt"

type Grid struct {
	Dimensions []GuessCompare
	GuessCount int
}

type GuessCompare struct {
	TargetValue   int
	Maximum       int
	Minimum       int
	ErrorMessage  string
	DimensionName string
	LowHint       string
	HighHint      string
}

func (gc GuessCompare) String() string {
	return fmt.Sprintf("This is the %v",gc.DimensionName)
}
/*
func (gc GuessCompare) Error() string {
	return gc.ErrorMessage
}
*/


func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func (g *Grid) ProcessGuess(raw_guess string) (GuessResult, error) {
	result := GuessResult{}
	
	valid_len := 3
	err_msg := "Guess must be in the format #,#,#.  Example:  5,5,5"

	parts := strings.Split(raw_guess, ",")
	if len(parts) != valid_len {
		err := errors.New(err_msg)
		return result, err
	}
	/*
	guess_x := 0
	guess_y := 0

	err_x := errors.New("")
	err_y := errors.New("")
	parts := strings.Split(raw_guess, ",")
	g.GuessCount += 1
	result.GuessCount = g.GuessCount
	if len(parts) != valid_len {
		err := errors.New(err_msg)
		return result, err
	}
	if guess_x, err_x = strconv.Atoi(parts[0]); err_x != nil {
		err := errors.New(err_msg)
		return result, err
	}

	if guess_y, err_y = strconv.Atoi(parts[1]); err_y != nil {
		err := errors.New(err_msg)
		return result, err
	}

	if guess_x < 1 || guess_y < 1 {
		err := errors.New("Coordinates must be greater than 0.")
		return result, err
	}
	err_outofbound := ""
	if guess_x > *width {
		err_outofbound = fmt.Sprintf("X coordinate can't be greater than %d.\n", *width)
	}
	if guess_y > *height {
		err_outofbound += fmt.Sprintf("Y coordinate can't be greater than %d.", *height)
	}
	if err_outofbound != "" {
		err := errors.New(err_outofbound)
		return result, err
	}

	if guess_x > g.TargetX {
		result.HorizontalPosition = cWest
	}
	if guess_x < g.TargetX {
		result.HorizontalPosition = cEast
	}
	if guess_x == g.TargetX {
		result.HorizontalPosition = cFound
	}

	if guess_y > g.TargetY {
		result.VerticalPosition = cNorth
	}
	if guess_y < g.TargetY {
		result.VerticalPosition = cSouth
	}
	if guess_y == g.TargetY {
		result.VerticalPosition = cFound
	}
	*/
	return result, nil
}

func (g *Grid) Build() {
	p := 10
	low := p / 2 * -1
	high := p / 2

	//Add ability to deal with odd numbers.
	//Anthing that is mod 2 !=0 needs to shift + or - one unit

	x := MakeGridCompare(low, high, "East", "West", "X Axis")
	y := MakeGridCompare(low, high, "North", "South", "Y Axis")
	z := MakeGridCompare(low, high, "Further", "Closer", "Z Axis")
}

func MakeGridCompare(min int, max int, lowhint string, highhint string, dimensionname string) GuessCompare {
	result := GuessCompare{}
	result.Minimum = min
	result.Maximum = max
	result.LowHint = lowhint
	result.HighHint = highhint
	result.DimensionName = dimensionname
	result.TargetValue = randInt(min, max)
	return result
}
