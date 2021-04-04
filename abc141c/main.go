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
	var n, k, q int
	fmt.Scanf("%d %d %d", &n, &k, &q)

	score := make([]int, n)
	for i := 0; i < q; i++ {
		score[nextInt()-1]++
	}

	for _, v := range score {
		if k-(q-v) > 0 {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
