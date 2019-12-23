package main

import (
	"fmt"
	"sort"
)

func main() {
	bell := make([]int, 3)
	for i := 0; i < 3; i++ {
		fmt.Scan(&bell[i])
	}	
	
	sort.Sort(sort.IntSlice(bell))

	fmt.Println(bell[0]+bell[1])
}