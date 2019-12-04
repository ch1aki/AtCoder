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

	var alice, bob int
	var flag bool = true
	for i := n - 1; 0 <= i; i-- {
		if flag {
			alice += cards[i]
			flag = false
		} else {
			bob += cards[i]
			flag = true
		}
	}

	fmt.Println(alice - bob)
}
