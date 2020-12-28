package main

import (
	"fmt"
	"math"
)

func main() {
	var n, w float64
	fmt.Scan(&n, &w)

	fmt.Println(math.Floor(n / w))

}
