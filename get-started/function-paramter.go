package main

import "fmt"

func sayHello(firstname string, lastname string) {
	fmt.Println("Hello ", firstname, lastname)
}

func main() {
	names := []string{"Agisna", "Revaldo"}
	sayHello(names[0], names[1])

	names = append(names, "Muhamad")
	sayHello(names[2], names[1])
}
