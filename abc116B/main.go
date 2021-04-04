package main

import "fmt"

func main() {
	var s int
	fmt.Scan(&s)

	a := make([]int, 1000000)
	a[0] = s

	r := 0
	for i := 1; r < 1; i++ {
		if a[i-1]%2 == 0 {
			a[i] = a[i-1] / 2
		} else {
			a[i] = 3*a[i-1] + 1
		}

		for j := 0; j < i; j++ {
			if a[j] == a[i] {
				r = i
				break
			}
		}
	}

	fmt.Println(r + 1)
}
