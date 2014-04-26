package main

const ( // iota is reset to 0
	cFound = iota
	cNorth = iota
	cSouth = iota
	cEast  = iota
	cWeat  = iota
)

type GuessResult struct {
	VerticalPosition   int
	HorizontalPosition int
	Message string
}


