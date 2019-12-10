package main

import (
	"fmt"
)

func main() {
	var n byte
	var s string
	fmt.Scan(&n)
	fmt.Scan(&s)

	sRune := []rune(s)

	for _, c := range sRune {
		rot := (byte(c)-'A'+n)%26 + 'A'
		fmt.Printf("%s", string(rot))
	}
}
