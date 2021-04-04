package main

import "fmt"

func main() {
	var n, k, x int
	fmt.Scan(&n, &k)

	r := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&x)

		a := x * 2
		b := (k - x) * 2

		if a < b {
			r += a
		} else {
			r += b
		}
	}

	fmt.Println(r)
}
