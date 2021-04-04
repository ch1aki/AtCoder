package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()

	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}

	r := -1
	max := 0

	for i := 2; i <= 1000; i++ {
		x := 0
		for j := 0; j < n; j++ {
			if a[j]%i == 0 {
				x++
			}
		}
		if max < x {
			max = x
			r = i
		}
	}

	fmt.Println(r)
}
