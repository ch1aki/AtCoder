package main

import "fmt"

func main() {
	var n int
	var s string

	fmt.Scan(&n)
	fmt.Scan(&s)

	x, max := 0, 0
	for _, c := range s {
		if c == 'I' {
			x++
		} else if c == 'D' {
			x--
		}

		if max < x {
			max = x
		}
	}

	fmt.Println(max)
}
