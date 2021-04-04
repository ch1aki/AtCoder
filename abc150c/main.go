package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"reflect"
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

func remove(ar []int, i int) []int {
	tmp := make([]int, len(ar))
	copy(tmp, ar)
	return append(tmp[0:i], tmp[i+1:]...)
}

func permutation(ar []int) [][]int {
	var result [][]int
	if len(ar) == 1 {
		return append(result, ar)
	}
	for i, a := range ar {
		for _, b := range permutation(remove(ar, i)) {
			result = append(result, append([]int{a}, b...))
		}
	}
	return result
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()

	p, q := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = nextInt()
	}
	for i := 0; i < n; i++ {
		q[i] = nextInt()
	}

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}

	indexP, indexQ := 0, 0
	for i, v := range permutation(nums) {
		if reflect.DeepEqual(v, p) {
			indexP = i
		}
		if reflect.DeepEqual(v, q) {
			indexQ = i
		}
	}

	fmt.Println(math.Abs(float64(indexP - indexQ)))
}
