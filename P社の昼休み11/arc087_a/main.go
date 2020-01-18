package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := map[int]int{}
	cnt := 0

	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)
		if a[x] < x {
			a[x]++
		} else {
			cnt++
		}
	}

	for k, v := range a {
		if k != v {
			cnt += v
		}
	}

	fmt.Println(cnt)
}
