package main

import "fmt"

func split(sum float32) (x, y float32) {
	x = sum * 1 / 3
	y = sum - x

	// return x, y (disebut "named result")
	return // disebut "naked result"
}

func main() {
	fmt.Println(split(10))
}
