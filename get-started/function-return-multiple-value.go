package main

import "fmt"

func getFullname(firstname string, lastname string) (string, string) {
	return firstname, lastname
}

func main() {
	firstname, lastname := getFullname("Agisna", "Revaldo")
	fmt.Println(firstname, lastname)

	username, _ := getFullname("Agisna", "Revaldo") // menggunakan blank identifier
	fmt.Println(username)
}
