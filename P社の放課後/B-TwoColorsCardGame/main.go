package main

import (
	"fmt"
)

func main() {
	var n, m int

	fmt.Scan(&n)
	blue := map[string]int{}
	for i := 0; i < n; i++ {
		var name string
		fmt.Scan(&name)
		blue[name]+=1
	}

	fmt.Scan(&m)
	red := map[string]int{}
	for i := 0; i < m; i++ {
		var name string
		fmt.Scan(&name)
		red[name]+=1
	}	

	var result int
	for k, v := range blue {
		_, exist := red[k]
		if exist {
			if v - red[k] > result {
				result = v - red[k]
			}
		} else {
			if v > result {
				result = v
			}
		}
	}

	fmt.Println(result)
}