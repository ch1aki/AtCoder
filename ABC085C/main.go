package main

import (
	"fmt"
)

func main() {
	var n, y int
	fmt.Scanf("%d %d\n", &n, &y)

	for t1 := 0; t1 <= n; t1++ {
		for t2 := 0; t2 <= n - t1; t2++ {
			t3 := n-t1-t2
			if 10000 * t1 + 5000 * t2 + 1000 * t3 == y {
				fmt.Printf("%d %d %d\n", t1, t2, t3)
				return
			}
		}
	}
	fmt.Println("-1 -1 -1")
}