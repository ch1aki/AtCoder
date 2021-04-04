package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	x := make([]int, n)
	y := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&x[i], &y[i])
	}

	r := 0
	for i := 0; i < n; i++ {
		for j := i+1; j < n; j++ {
			if i != j {
				tilt := float64(y[j] - y[i]) / float64(x[j] - x[i])
				if tilt >= -1 && tilt <= 1 {
					r += 1
				}
			}
		}
	}

	fmt.Println(r)
}
