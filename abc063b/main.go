package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	result := true
	m := make(map[rune]int)
	for _, c := range s {
		if m[c] == 0 {
			m[c]++
		} else {
			result = false
			break
		}
	}

	if result {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
