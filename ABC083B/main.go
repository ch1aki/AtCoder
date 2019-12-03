package main

import "fmt"

func main() {
	var n, a, b int
	var r int
	fmt.Scanf("%d %d %d", &n, &a, &b)

	for i := 1; i <= n; i++ {
		var sum int
		var num int = i

		for num >= 1 {
			sum += num % 10
			num = num / 10
		}

		if a <= sum && sum <= b {
			r += i
		}
	}
	fmt.Println(r)
}
