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
	sort.Ints(cards)

	var r int
	op := 1

	for i := n - 1; 0 <= i; i-- {
		r += cards[i] * op
		op *= -1
	}

	fmt.Println(r)
}
