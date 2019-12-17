package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, x int
	fmt.Scanf("%d %d %d", &a, &b, &x)

	start := 0
	end := int(math.Pow(10, 9))
	var r int
	for {
		if end < start {
			break
		}

		n := (start + end)/2
		t := a * n + b * int(math.Log10(float64(n)))
		fmt.Printf("start:%d\tent:%d\tn:%d\tr:%d\tt:%d\n", start, end, n, r, t) // DEBUG

		if x >= t && t > r {
			r = n
		}

		if x > t {
			start = n
		} else {
			end = n
		}
	}

	fmt.Println(r)

}
