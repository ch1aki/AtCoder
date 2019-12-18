package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b, x int
	fmt.Scanf("%d %d %d", &a, &b, &x)

	start := 0
	end := int(1e9) + 1

	for (end - start) > 1 {
		n := (start + end)/2

		t := a * n + b * len(strconv.Itoa(n))

		if t <= x {
			start = n
		} else {
			end = n
		}
	}

	fmt.Println(start)

}
