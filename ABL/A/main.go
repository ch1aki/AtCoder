package main

import "fmt"

func main() {
	var k int
	fmt.Scan(&k)
	r := ""
	acl := "ACL"

	for i := 0; i < k; i++ {
		r = r + acl
	}

	fmt.Println(r)
}
