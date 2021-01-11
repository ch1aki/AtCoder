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
	var n, m, c int
	fmt.Scan(&n, &m, &c)

	b := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&b[i])
	}

	sc.Split(bufio.ScanWords)

	r := 0
	for i := 0; i < n; i++ {
		tmp := c
		for j := 0; j < m; j++ {
			tmp += nextInt() * b[j]
		}
		if tmp > 0 {
			r++
		}
	}

	fmt.Println(r)
}
