package main

import (
	"fmt"
)

func main() {
	var oRaw, eRaw string
	fmt.Scan(&oRaw)
	fmt.Scan(&eRaw)

	o := []rune(oRaw)
	e := []rune(eRaw)

	var limit int
	if len(o) > len(e) {
		limit = len(e)
	} else {
		limit = len(o)	
	}
	for i := 0; i < limit; i++ {
		fmt.Print(string(o[i]))
		fmt.Print(string(e[i]))
	}

	if len(e) - len(o) == 1 {
		fmt.Print(string(e[len(e)-1]))
	}
	if len(o) - len(e) == 1 {
		fmt.Print(string(o[len(o)-1]))
	}
}