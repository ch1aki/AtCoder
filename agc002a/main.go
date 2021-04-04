package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)

	if 0 < a {
		fmt.Println("Positive")
	} else if a <= 0 && 0 <= b {
		fmt.Println("Zero")
	} else if b < 0 {
		if (b-a)%2 == 0 {
			fmt.Println("Negative")
		} else {
			fmt.Println("Positive")
		}
	}
}
