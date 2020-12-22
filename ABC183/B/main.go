package main

import "fmt"

func main() {
	var sx, sy, gx, gy float64
	fmt.Scan(&sx, &sy, &gx, &gy)

	ratio := sy / (gy + sy)

	fmt.Println(sx + (gx-sx)*ratio)
}
