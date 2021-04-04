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
	sc.Split(bufio.ScanWords)
	n := nextInt()

	v := make([]float64, n)
	for i := 0; i < n; i++ {
		v[i] = float64(nextInt())
	}

	for 2 <= len(v) {
		sort.Sort(sort.Float64Slice(v))
		x, y := v[0], v[1]
		v = v[2:]
		v = append(v, (x+y)/2)
	}

	fmt.Println(v[0])
}
