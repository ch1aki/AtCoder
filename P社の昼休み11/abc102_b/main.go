package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	min := a[0]
	max := a[0]

	for _, i := range a {
		if i < min {
			min = i
		} else if i > max {
			max = i
		}
	}

	fmt.Println(max - min)
}
