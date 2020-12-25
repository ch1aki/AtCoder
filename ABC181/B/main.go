package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	r := 0

	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		r += (b + a) * (b - a + 1) / 2
	}

	fmt.Println(r)
}
