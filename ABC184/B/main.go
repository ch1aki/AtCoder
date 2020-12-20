package main

import "fmt"

func main() {
	var n, x int
	var s string

	fmt.Scan(&n, &x)
	fmt.Scan(&s)

	for _, i := range s {
		if i == 'o' {
			x = x + 1
		} else if x > 0 {
			x = x - 1
		}
	}

	fmt.Println(x)
}
