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

	n, m := nextInt(), nextInt()
	a := make(map[int]int, m)
	for i := 0; i < n; i++ {
		k := nextInt()
		for j := 0; j < k; j++ {
			a[nextInt()]++
		}
	}

	r := 0
	for _, v := range a {
		if v == n {
			r++
		}
	}

	fmt.Println(r)
}
