package main

import "fmt"
import "math/rand"

type Point struct {
	x       int
	y       int
	target  bool
	explore bool
}

func (p Point) String() string {
	return fmt.Sprintf("{X:%d, Y:%d - Target: %t, Explored %t}", p.x, p.y, p.target, p.explore)
}

type Grid struct {
	Points []Point
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func (g *Grid) Build() {
	target_x := randInt(1, *width)
	target_y := randInt(1, *height)

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
	fmt.Printf("Target (X,Y) (%d,%d)", target_x, target_y)
}
