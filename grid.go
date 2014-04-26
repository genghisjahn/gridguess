package main

import "fmt"
import "math/rand"
import "strings"
import "errors"
import "strconv"

type Grid struct {
	Points  []Point
	TargetX int
	TargetY int
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func (g *Grid) IsPointTarget(x int, y int) bool {
	result := false

	for _, point := range g.Points {
		if point.x == x && point.y == y {
			point.explored = true
			if point.target {
				result = true
				break
			}
		}
	}
	return result
}

func (g *Grid) ProcessGuess(raw_guess string) (GuessResult, error) {
	result := GuessResult{}
	err_msg := "Guess must be in the format #,#.  Example:  5.5"
	guess_x:=0
	guess_y:=0
	err_x := errors.New("")
	err_y := errors.New("")
	parts := strings.Split(raw_guess, ",")
	if len(parts) != 2 {
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
	guess_x+=1
	guess_y+=1
	return result, nil
}

func (g *Grid) Build() {
	target_x := randInt(1, *width)
	target_y := randInt(1, *height)

	g.TargetX = target_x
	g.TargetY = target_y

	set_x := 1
	set_y := 1
	for set_x <= *width {
		for set_y <= *height {
			tempPoint := Point{}
			tempPoint.x = set_x
			tempPoint.y = set_y
			set_y += 1

			if tempPoint.x == target_x && tempPoint.y == target_y {
				tempPoint.target = true
			}

			g.Points = append(g.Points, tempPoint)
			fmt.Printf("Here is the point %v.\n", tempPoint)
		}
		set_x += 1
		set_y = 1
	}

	gresult := GuessResult{}
	gresult.VerticalPosition = cNorth
}
