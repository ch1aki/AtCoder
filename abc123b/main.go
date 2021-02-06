package main

import "fmt"

func Permute(nums []int) [][]int {
	n := factorial(len(nums))
	ret := make([][]int, 0, n)
	permute(nums, &ret)
	return ret
}

func permute(nums []int, ret *[][]int) {
	*ret = append(*ret, makeCopy(nums))

	n := len(nums)
	p := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		p[i] = i
	}
	for i := 1; i < n; {
		p[i]--
		j := 0
		if i%2 == 1 {
			j = p[i]
		}

		nums[i], nums[j] = nums[j], nums[i]
		*ret = append(*ret, makeCopy(nums))
		for i = 1; p[i] == 0; i++ {
			p[i] = i
		}
	}

}

func factorial(n int) int {
	ret := 1
	for i := 2; i <= n; i++ {
		ret *= i
	}
	return ret
}

func makeCopy(nums []int) []int {
	return append([]int{}, nums...)
}

func slove(foods []int) int {
	r := 0
	for i := 0; i < len(foods)-1; i++ {
		if foods[i]%10 > 0 {
			r += foods[i] + (10 - foods[i]%10)
		} else {
			r += foods[i]
		}
	}

	r += foods[len(foods)-1]
	return r
}

func main() {
	foods := make([]int, 5)
	for i := 0; i < 5; i++ {
		fmt.Scan(&foods[i])
	}

	min := 643
	for _, v := range Permute(foods) {
		r := slove(v)
		if r < min {
			min = r
		}
	}

	fmt.Println(min)
}
