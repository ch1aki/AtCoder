package main

import "fmt"

func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)

	var lMax, rMin int
	fmt.Scanf("%d %d", &lMax, &rMin)

	for i := 1; i < m; i++ {
		var l, r int
		fmt.Scanf("%d %d", &l, &r)

		if l > lMax {
			lMax = l
		}

		if r < rMin {
			rMin = r
		}
	}

	r := 0
	for i := lMax; i <= rMin; i++ {
		r++
	}

	fmt.Println(r)
}
