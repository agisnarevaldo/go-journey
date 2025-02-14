package main

import "fmt"

func multiple(x int, y int) int {
	return x * y
}

func main() {
	var (
		a = 8
		b = 9
	)

	fmt.Println(multiple(a, b))
}
