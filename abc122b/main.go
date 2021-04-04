package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	r := 0
	for i := 0; i < len(s); i++ {
		cnt := 0
		for _, c := range s[i:] {
			if c == 'A' || c == 'C' || c == 'G' || c == 'T' {
				cnt++
				if cnt > r {
					r = cnt
				}
			} else {
				break
			}
		}
	}
	fmt.Println(r)
}
