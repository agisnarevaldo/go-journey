package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func sayHell() {
	fmt.Println("Hello")
}

func main() {
	sayHell()
	fmt.Println(add(10, 12))
}
