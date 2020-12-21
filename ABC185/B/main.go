package main

import "fmt"

func scan(n int, m int, t int) bool {
	last_time := 0
	current := n

	for i := 0; i < m; i++ {
		var a, b int
		fmt.Scan(&a, &b)

		// drain
		current -= a - last_time
		if current <= 0 {
			return false
		}
		last_time = b

		// charge
		charge := b - a
		if charge+current > n {
			current = n
		} else {
			current += charge
		}
	}

	if t-last_time < current {
		return true
	}
	return false
}

func main() {
	var n, m, t int
	fmt.Scan(&n, &m, &t)

	if scan(n, m, t) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
