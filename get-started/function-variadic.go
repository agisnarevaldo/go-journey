package main

import "fmt"

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func multiple(numbers ...int) int {
	total := 1
	for _, number := range numbers {
		total *= number
	}
	return total
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(multiple(1, 2, 3, 4, 5))

	numbers := []int{10, 10, 10, 10, 10}
	fmt.Println(sum(numbers...))
	fmt.Printf("type: %T", numbers)
}
