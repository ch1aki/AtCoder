package main

import "fmt"

func attack(h int) int {
	if h == 1 {
		return 1
	} else {
		return 2*attack(h/2) + 1
	}
}

func main() {
	var h int
	fmt.Scan(&h)

	fmt.Println(attack(h))
}
