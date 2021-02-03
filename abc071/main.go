package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	occur := make([]bool, 26)
	for _, c := range s {
		occur[c-97] = true
	}

	r := "None"
	for i, v := range occur {
		if v == false {
			r = string('a' + i)
			break
		}
	}

	fmt.Println(r)
}
