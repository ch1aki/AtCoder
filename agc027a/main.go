package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	var x, n, cnt int
	fmt.Scan(&n, &x)

	a := make([]int, n)

	sc.Split(bufio.ScanWords)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}

	sort.Ints(a)
	for i := 0; i < n; i++ {
		if a[i] <= x {
			cnt++
		}
		x -= a[i]
	}
	if 0 < cnt && 0 < x {
		cnt--
	}

	fmt.Println(cnt)
}
