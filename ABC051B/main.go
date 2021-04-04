package main

import "fmt"

func main() {
	var r, k, s int
	fmt.Scan(&k, &s)

	xMax := s
	if k < s {
		xMax = k
	}
	for x := 0; x <= xMax; x++ {
		yMax := s - x
		if k < yMax {
			yMax = k
		}
		for y := 0; y <= yMax; y++ {
			if s-x-y <= k {
				r++
			}
		}
	}

	fmt.Println(r)
}
