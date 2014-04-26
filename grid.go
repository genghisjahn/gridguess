package main

import "fmt"
// import "math/rand"


type Point struct{
	x int
	y int
	target bool
	explore bool
}

func (p Point) String() string {
    return fmt.Sprintf("{X:%d, Y:%d - Target: %t, Explored %t}", p.x, p.y,p.target,p.explore)
}

type Grid struct{
	Points []Point
}


