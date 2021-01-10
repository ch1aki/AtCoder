package main

import (
	"fmt"
	"math"
)

func main() {
	var n float64
	fmt.Scan(&n)

	r := 0
	for x := 1; x <= 50000; x++ {
		if n == math.Floor(float64(x)*1.08) {
			r = x
			break
		}
	}

	if r > 0 {
		fmt.Println(r)
	} else {
		fmt.Println(":(")
	}
}
