package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, x int
	fmt.Scanf("%d %d %d", &a, &b, &x)

	n := math.Pow(10, 9)
	for x <= (a*int(n)+b*int(math.Log10(n))) && n > 0 {
		n--
	}

	fmt.Println(n)

}
