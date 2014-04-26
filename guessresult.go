package main

import "fmt"

const (
	cInvalid = iota
	cNorth = iota
	cSouth = iota
	cEast  = iota
	cWest  = iota
	cFound = iota
)

type GuessResult struct {
	VerticalPosition   int
	HorizontalPosition int
	GuessCount         int
}

func (gr GuessResult) String() string {
	v_part := "Invalid"
	h_part := "Invalid"
	if gr.VerticalPosition == cNorth {
		v_part = "North"
	}
	if gr.VerticalPosition == cSouth {
		v_part = "South"
	}
	if gr.VerticalPosition == cFound {
		v_part = "Found"
	}

	if gr.HorizontalPosition == cEast {
		h_part = "East"
	}
	if gr.HorizontalPosition == cWest {
		h_part = "West"
	}
	if gr.HorizontalPosition == cFound {
		h_part = "Found"
	}
	return fmt.Sprintf("Guess #%v: Target is %v and %v from your guess.",gr.GuessCount, v_part, h_part)
}
