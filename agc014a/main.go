package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	r := 0
	for i := 0; a%2 == 0 && b%2 == 0 && c%2 == 0; i++ {
		nextA, nextB, nextC := b/2+c/2, a/2+c/2, a/2+b/2
		if a == nextA && b == nextB && c == nextC {
			r = -1
			break
		}

		a, b, c = nextA, nextB, nextC
		r++
	}

	fmt.Println(r)
}
