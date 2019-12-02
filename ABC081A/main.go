package main

import "fmt"

func main() {
	var in string
	var cnt int
	fmt.Scanln(&in)
	for _, s := range in {
		if s == '1' {
			cnt++
		}
	}
	fmt.Println(cnt)
}