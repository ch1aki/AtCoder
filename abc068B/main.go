package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	result, countMax := 1, 0
	for i := 1; i <= n; i++ {
		x := i
		cnt := 0
		for j := 0; x%2 == 0; j++ {
			x /= 2
			cnt++
		}

		if cnt > countMax {
			result = i
			countMax = cnt
		}
	}

	fmt.Println(result)
}
