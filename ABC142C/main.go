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

	a := make(map[int]int, n)
	for i := 1; i <= n; i++ {
		a[nextInt()] = i
	}

	for i := 1; i <= n; i++ {
		fmt.Printf("%d ", a[i])
	}
}
