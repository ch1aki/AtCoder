package main

import (
	"fmt"
)

func main() {
	var count int
	coins := make([]int, 3)
	for i := 0; i < 3; i++ {
		fmt.Scan(&coins[i])
	}

	var x int
	fmt.Scan(&x)

	for i := 0; i <= coins[2]; i++ {
		for j := 0; j <= coins[1]; j++ {
			for k := 0; k <= coins[0]; k++ {
				if ((x-50*i)-100*j)-500*k == 0 {
					count++
				}
			}
		}
	}
	fmt.Println(count)
}