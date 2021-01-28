package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	var input string
	fmt.Scan(&input)

	min := 0.0
	s := strings.Split(input, "")
	for i := 0; i < len(s)-2; i++ {
		x, _ := strconv.Atoi(strings.Join(s[i:i+3], ""))
		abs := math.Abs(753 - float64(x))
		if i == 0 || min > abs {
			min = abs
		}
	}

	fmt.Println(min)
}
