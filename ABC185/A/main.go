package main

import "fmt"

func main() {
	var min int

	for i := 0; i < 4; i++ {
		var a int
		fmt.Scan(&a)

		if min == 0 || a < min {
			min = a
		}
	}
	fmt.Println(min)
}
