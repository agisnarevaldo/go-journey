package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Muhamad"
	names[1] = "Agisna"
	names[2] = "Revaldo"

	fmt.Println(names)

	var values = [...]int{
		1,
		2,
		3,
	}
	fmt.Println(values)
	fmt.Printf("%T\n", values)
	fmt.Println(len(values))
}
