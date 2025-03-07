package main

import "fmt"

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func addLoop(value int) int {
	result := 0
	for i := value; i < 10; i++ {
		result += i
	}
	return result
}

func addRecursive(value int) int {
	if value == 10 {
		return 0
	} else {
		return value + addRecursive(value+1)
	}
}

func main() {
	fmt.Println(factorialLoop(10))
	fmt.Println(factorialRecursive(10))
	fmt.Println(addLoop(1))
	fmt.Println(addRecursive(1))
}
