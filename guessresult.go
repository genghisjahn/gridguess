package main

import "fmt"

const ( // iota is reset to 0
	cFound = iota
	cNorth = iota
	cSouth = iota
	cEast  = iota
	cWest  = iota
)

type GuessResult struct {
	VerticalPosition   int
	HorizontalPosition int
	Message            string
}

func (gr GuessResult) String() string {
	return fmt.Sprintf("Still working on this!!")
}

