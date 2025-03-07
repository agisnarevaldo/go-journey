package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	fmt.Println("Hello", filter(name))
}

func addSpam(word string) string {
	if word == "Fuck" {
		word = "***"
	}
	return word
}

func addLimit(word string) string {
	if len(word) > 10 {
		word = "....."
	}
	return word
}

func main() {
	sayHelloWithFilter("agisna", addSpam)
	sayHelloWithFilter("Fuck", addSpam)
	sayHelloWithFilter("Muhamad Agisna Revaldo", addLimit)
}
