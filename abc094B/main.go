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
	var n, m, x int
	fmt.Scan(&n, &m, &x)

	sc.Split(bufio.ScanWords)

	cnt := 0
	for i := 0; i < m; i++ {
		a := nextInt()
		if a < x {
			cnt++
		}
	}

	if cnt < m-cnt {
		fmt.Println(cnt)
	} else {
		fmt.Println(m - cnt)
	}
}
