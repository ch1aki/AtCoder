package main

import (
	"fmt"
	"math"
)

func square(n float64) float64 {
	return n * n
}

func length(x1 float64, y1 float64, x2 float64, y2 float64) float64 {
	return math.Sqrt(square(x2-x1) + square(y2-y1))
}

func main() {
	var n int
	fmt.Scan(&n)

	p := make([][]float64, n)
	for i := 0; i < n; i++ {
		var x, y float64
		fmt.Scan(&x, &y)
		p[i] = append(p[i], x, y)
	}

	max := -1.0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			tmp := length(p[i][0], p[i][1], p[j][0], p[j][1])
			if max < tmp {
				max = tmp
			}
		}
	}

	fmt.Println(max)
}
