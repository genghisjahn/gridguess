package main

import "math/rand"
import "strings"
import "errors"
import "strconv"
import "fmt"

type Grid struct {
	TargetX    int
	TargetY    int
	GuessCount int
}

type GuessCompare struct {
	IsActive bool
	Value int
	Maximum int
	ErrorMessage string
	DimensionName string
}
func (gc GuessCompare) Error() string{
	return gc.ErrorMessage
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func (g *Grid) ProcessGuess(raw_guess string) (GuessResult, error) {
	result := GuessResult{}
	valid_len := 2
	err_msg := "Guess must be in the format #,#.  Example:  5.5"
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
	return result, nil
}

func (g *Grid) Build() {
	target_x := randInt(1, *width)
	target_y := randInt(1, *height)

	g.TargetX = target_x
	g.TargetY = target_y

}
