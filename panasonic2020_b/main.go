package main

import (
	"fmt"
	"math"
)

func main() {
	var h, w float64
	fmt.Scan(&h, &w)

	if h == 1 || w == 1 {
		fmt.Println(1)
	} else {
		fmt.Println(int(math.Ceil(h * w / 2)))
	}

}
