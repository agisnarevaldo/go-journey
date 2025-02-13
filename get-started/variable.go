package main

import "fmt"

func main() {
	// definisi variable dengan tipe data-nya
	var name string

	name = "Muhamad Agisna"
	fmt.Println(name)

	name = "Muhamad Revaldo"
	fmt.Println(name)

	// definisi variable dengan value-nya (tidak harus menyertakan tipe data)
	age := 22
	fmt.Println(age)

	age = 23
	fmt.Println(age)

	// definisi banyak variable
	var (
		firstname  = "muhamad"
		middlename = "Agisna"
		lastname   = "revaldo"
	)

	fmt.Println(firstname)
	fmt.Println(middlename)
	fmt.Println(lastname)
}
