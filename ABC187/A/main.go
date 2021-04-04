package main

import (
	"fmt"
	"strconv"
)

func sum(n int) int {
	r := 0
	for _, v := range strconv.Itoa(n) {
		r += int(v - '0')
	}
	return r
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	sumA := sum(a)
	sumB := sum(b)

	if sumA > sumB {
		fmt.Println(sumA)
	} else {
		fmt.Print(sumB)
	}
}
