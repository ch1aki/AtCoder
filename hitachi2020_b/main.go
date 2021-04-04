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
	var a, b, m int
	fmt.Scan(&a, &b, &m)

	as, bs := make([]int, a), make([]int, b)
	aMin, bMin := 0, 0
	sc.Split(bufio.ScanWords)
	for i := 0; i < a; i++ {
		as[i] = nextInt()
		if i == 0 || aMin > as[i] {
			aMin = as[i]
		}
	}
	for i := 0; i < b; i++ {
		bs[i] = nextInt()
		if i == 0 || bMin > bs[i] {
			bMin = bs[i]
		}
	}

	min := aMin + bMin

	for i := 0; i < m; i++ {
		x, y, c := nextInt(), nextInt(), nextInt()

		total := as[x-1] + bs[y-1] - c
		if min > total {
			min = total
		}
	}

	fmt.Println(min)
}
