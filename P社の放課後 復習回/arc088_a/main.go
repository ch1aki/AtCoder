package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Scanf("%d %d", &x, &y)

	cnt := 0
	for x <= y {
		cnt++
		x *= 2
	}

	fmt.Println(cnt)
}
