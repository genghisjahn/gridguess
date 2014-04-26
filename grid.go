package main

import "fmt"
// import "math/rand"


type Point struct{
	x int
	y int
	target bool
}

func (p Point) String() string {
    return fmt.Sprintf("{X:%d, Y:%d - Target: %t}", p.x, p.y,p.target)
}



