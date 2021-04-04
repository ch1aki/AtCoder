package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	if strings.HasSuffix(s, "s") {
		s += "es"
	} else {
		s += "s"
	}

	fmt.Println(s)
}
