package main

import "fmt"

func main() {
	name := "Revaldo"

	if name == "Agisna" {
		fmt.Printf("Hello, how are you %v?\n", name)
	} else if name == "Revaldo" {
		fmt.Printf("Hello, how old are you %v\n", name)
	} else {
		fmt.Printf("Hello, are you %v?\n", name)
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama lebih dari 5 karakter")
	} else {
		fmt.Println("Nama sudah pas")
	}
}
