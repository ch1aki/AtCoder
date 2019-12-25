package main

import (
	"fmt"
)

func main() {
	var n int
	var s string
	fmt.Scan(&n)
	fmt.Scan(&s)

	c := n/2
	if s[0:c] == s[c:n] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}