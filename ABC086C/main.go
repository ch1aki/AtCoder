package main

import (
	"fmt"
	"math"
)

type Point struct {
	T int
	X int
	Y int
}

type Points []Point

func main() {
	var points Points
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		p := Point {}
		fmt.Scanf("%d %d %d\n", &p.T, &p.X, &p.Y)
		points = append(points, p)
	}

	c := Point{0, 0, 0}
	for _, p := range(points) {
		moving := p.T - c.T
		dist := int(math.Abs(float64(p.X - c.X)) + math.Abs(float64(p.Y - c.Y)))
		if moving < dist {
			fmt.Println("No")
			return
		} else if moving > dist && moving % 2 != dist % 2 {
			fmt.Println("No")
			return
		}
		c = p
	}
	fmt.Println("Yes")
}