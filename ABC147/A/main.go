package main

import (
	"fmt"
)

func main() {
	var a1, a2, a3 int
	fmt.Scanf("%d %d %d", &a1, &a2, &a3)

	if a1 + a2 + a3 >= 22 {
		fmt.Println("bust")
	} else {
		fmt.Println("win")
	}
}