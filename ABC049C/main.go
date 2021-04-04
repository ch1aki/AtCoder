package main

import (
	"fmt"
)

func main() {
	var s, t string
	fmt.Scan(&s)

	fmt.Printf("input: %s\n", s) // DEBUG

	matches := [4]string{
		"dreamer",
		"eraser",
		"dream",
		"erace",
	}

	start := 0
	for {
		for _, m := range matches {
			fmt.Printf("DEBUG: %v\n", s[start:len(m)]) // DEBUG
			if start+len(m) < len(s) && s[start:len(m)-1] == m {
				t = t + m
				fmt.Printf("t: %s\n", t) // DEBUG
				start += len(m)
				break
			}
		}
		if start > len(s) - 4 {
			break
		}
	}

	fmt.Printf("t: %s\n", t) // DEBUG
	if s == t {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

