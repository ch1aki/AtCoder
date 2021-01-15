package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scanln(&n)

	cards := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&cards[i])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(cards)))

	a, b := 0, 0
	for i, v := range cards {
		if i%2 == 0 {
			a += v
		} else {
			b += v
		}
	}

	fmt.Println(a - b)
}
