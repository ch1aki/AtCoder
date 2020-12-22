package main

import "fmt"

func main() {
	var x, r int
	fmt.Scan(&x)

	if x >= 0 {
		r = x
	} else {
		r = 0
	}

	fmt.Println(r)
}
