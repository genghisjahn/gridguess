package main

import "fmt"

type Point struct {
	x        int
	y        int
	target   bool
	explored bool
}

func (p Point) String() string {
	return fmt.Sprintf("{X:%d, Y:%d - Target: %t, Explored %t}", p.x, p.y, p.target, p.explored)
}
