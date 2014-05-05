package main

import "math"
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
	Description   string
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

func (g *Grid) Build(length int) {
	low := 0
	high := 0
	if math.Mod(float64(length), 2) == 0 {
		low = length / 2 * -1
		high = length / 2
	} else {
		low = (length - 1) / 2 * -1
		high = (length + 1) / 2
	}

	x_desc = fmt.Scanf(" Max West: %v.   Max East: %v.\n", low, high)
	y_desc = fmt.Scanf("Max South: %v.   Max North: %v.\n", low, high)
	z_desc = fmt.Scanf(" Furthest: %v.     Closest: %v.\n", low, high)

	x := MakeGridDimension(low, high, "East", "West", "X Axis", x_desc)
	y := MakeGridDimension(low, high, "North", "South", "Y Axis", y_desc)
	z := MakeGridDimension(low, high, "Closer", "Further", "Z Axis", z_desc)
	g.Dimensions = append(g.Dimensions, x, y, z)
	g.GuessCount = 1
}

func MakeGridDimension(min int, max int, lowhint string, highhint string, dimensionname string, description string) Dimension {
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
