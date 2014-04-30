package main

import "math/rand"
import "strings"
import "errors"
import "strconv"
import "fmt"

type Grid struct {
	Dimensions []Dimension
	GuessCount int
}

type Dimension struct {
	TargetValue   int
	Maximum       int
	Minimum       int
	ErrorMessage  string
	DimensionName string
	LowHint       string
	HighHint      string
	Found         bool
}

func (gc Dimension) String() string {
	return fmt.Sprintf("This is the %v", gc.DimensionName)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func (g *Grid) ProcessGuess(raw_guess string) (GuessResult, error) {
	result := GuessResult{}
	g.GuessCount += 1
	valid_len := 3
	err_msg := "Guess must be in the format #,#,#.  Example:  5,5,5"

	parts := strings.Split(raw_guess, ",")
	if len(parts) != valid_len {
		err := errors.New(err_msg)
		return result, err
	}
	guess_coordinates := make([]int, 3)

	for index, g := range parts {
		guess_val := 0
		err_d := errors.New("")
		guess_val, err_d = strconv.Atoi(g)
		if err_d != nil {
			err := errors.New(err_msg)
			return result, err
		}
		guess_coordinates[index] = guess_val
	}
	temp := ""
	result.Found = true
	for index, dimension := range g.Dimensions {
		if guess_coordinates[index] < dimension.TargetValue {
			temp += " - " + dimension.LowHint
			result.Found = false
		}
		if guess_coordinates[index] > dimension.TargetValue {
			temp += " - " + dimension.HighHint
			result.Found = false
		}
		if guess_coordinates[index] == dimension.TargetValue {
			temp += " - " + dimension.DimensionName + " is correct."
		}
	}
	if result.Found {
		result.Message = "You guessed the point!"
	} else {
		result.Message = temp
	}

	return result, nil
}

func (g *Grid) Build() {
	p := 10
	low := p / 2 * -1
	high := p / 2

	//Add ability to deal with odd numbers.
	//Anthing that is mod 2 !=0 needs to shift + or - one unit

	x := MakeGridDimension(low, high, "East", "West", "X Axis")
	y := MakeGridDimension(low, high, "North", "South", "Y Axis")
	z := MakeGridDimension(low, high, "Closer", "Further", "Z Axis")
	g.Dimensions = append(g.Dimensions, x, y, z)
	g.GuessCount = 1
}

func MakeGridDimension(min int, max int, lowhint string, highhint string, dimensionname string) Dimension {
	result := Dimension{}
	result.Minimum = min
	result.Maximum = max
	result.LowHint = lowhint
	result.HighHint = highhint
	result.DimensionName = dimensionname
	result.TargetValue = randInt(min, max)
	result.Found = false
	return result
}
