package main

import (
	"fmt"
)

func main() {
	var t, n, k, m int
	fmt.Scanf("%d %d %d", &n, &k, &m)

	for i := 0; i < n-1; i++ {
		var a int
		fmt.Scan(&a)
		t += a
	}

	x := m*n - t
	if x > k {
		fmt.Println(-1)
	} else if x < 0 {
		fmt.Println(0)
	} else {
		fmt.Println(x)
	}
}
