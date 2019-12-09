package main

import (
	"fmt"
)

type Point struct {
	T int
	X int
	Y int
}

type Points []Point

func distance(current Point, target)

func main() {
	var points Points
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		p := Point {}
		fmt.Scanf("%d %d %d\n", &p.T, &p.X, &p.Y)
		points = append(points, p)
	}

	fmt.Printf("%v\n", points) // DEBUG

}