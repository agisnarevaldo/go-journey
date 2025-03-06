package main

import "fmt"

func completedName() (firstName, middleName, lastName string) {
	firstName = "Muhamad"
	middleName = "Agisna"
	lastName = "Revaldo"

	return firstName, middleName, lastName
}

func main() {
	a, b, c := completedName()
	fmt.Println(a, b, c)
}
