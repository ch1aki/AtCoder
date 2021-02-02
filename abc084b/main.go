package main

import (
	"fmt"
	"regexp"
)

func main() {
	var a, b int
	var s string

	fmt.Scanf("%d %d", &a, &b)
	fmt.Scan(&s)

	reg := fmt.Sprintf("[0-9]{%d}-[0-9]{%d}", a, b)
	if regexp.MustCompile(reg).Match([]byte(s)) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
