package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println(swap("a", "b"))

	firstname, lastname := swap("agisna", "revaldo")
	fmt.Println(firstname, lastname)

	a := "a"
	b := "b"
	fmt.Println(swap(a, b))

	var (
		cur1 = "IDR"
		cur2 = "USD"
	)
	fmt.Println(swap(cur1, cur2))
}
