package main

import "fmt"

func getHello(name string) string {
	hello := "Hello " + name
	return hello
}

func main() {
	greetings := getHello("Agisna")
	fmt.Println(greetings)

	fmt.Println(getHello("Example"))
}
