package main

import (
	"bufio"
	"fmt"
	"math"
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

	var manhttan, euclid, chebyshev float64

	for i := 0; i < n; i++ {
		x := math.Abs(float64(nextInt()))

		manhttan += x
		euclid += x * x

		if i == 0 || chebyshev < x {
			chebyshev = x
		}
	}

	fmt.Println(manhttan)
	fmt.Println(math.Sqrt(euclid))
	fmt.Println(chebyshev)

}
