package main

import "fmt"

func slove(a, b, c int) bool {
	return a+b > c && b+c > a && c+a > b
}

func main() {
	var n int
	fmt.Scan(&n)

	l := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&l[i])
	}

	r := 0
	for a := 0; a < n; a++ {
		for b := 0; b < n; b++ {
			for c := 0; c < n; c++ {
				if a != b && b != c && l[a]+l[b] > l[c] && l[b]+l[c] > l[a] && l[c]+l[a] > l[b] {
					r++
				}
			}
		}
	}

	fmt.Println(r)
}
