package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	for i:=0;; i++ {
		if a * i - (i-1) >= b {
			fmt.Println(i)
			break
		}
	}
}
