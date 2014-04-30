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
}

func (gc Dimension) String() string {
	return fmt.Sprintf("This is the %v", gc.DimensionName)
}

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
	guess_coordinates := make([]int, 3)
	for _, g := range parts {
		guess := 0
		err_d := errors.New("")
		if guess, err_d = strconv.Atoi(g); err_d != nil {
			err := errors.New(err_msg)
			return result, err
		}
		guess_coordinates = append(guess_coordinates, guess)
	}

	/* Guess coordinates must be evaluted in the same order
	as the Dimensions on the Gris.
	*/
	temp := ""
	for index, dimension := range g.Dimensions {
		if guess_coordinates[index] < dimension.TargetValue {			
			temp += " - " + dimension.LowHint  
		}
		if guess_coordinates[index] > dimension.TargetValue {
			temp += " - " + dimension.HighHint  
		}
		if guess_coordinates[index] == dimension.TargetValue {
			temp += " - " + dimension.DimensionName + " is correct."  
		}
		fmt.Printf("Target %v: %v\n",dimension.DimensionName,dimension.TargetValue)
	}
	fmt.Printf(temp + "\n")
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
	z := MakeGridDimension(low, high, "Further", "Closer", "Z Axis")
	g.Dimensions = append(g.Dimensions, x, y, z)

	for _,d := range g.Dimensions{
		fmt.Printf("%v - %v.\n", d.DimensionName,d.TargetValue)
	}
}

func MakeGridDimension(min int, max int, lowhint string, highhint string, dimensionname string) Dimension {
	result := Dimension{}
	result.Minimum = min
	result.Maximum = max
	result.LowHint = lowhint
	result.HighHint = highhint
	result.DimensionName = dimensionname
	result.TargetValue = randInt(min, max)
	return result
}
