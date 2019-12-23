package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	sides := map[int]int{}
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		sides[a]+=1
	}	

	var first, second int
	for k, v := range sides {
		if k > first {
			if v >= 2 {
				second = first
				first = k
				v -= 2
			}
		}
		if k > second {
			if v >= 2 {
				second = k
			}
		}
	}

	fmt.Println(first * second)
}