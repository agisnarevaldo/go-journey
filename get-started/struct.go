package main

import "fmt"

type Customer struct {
	username, fullname string
	age                int
}

func (customer Customer) sayhello(name string) {
	fmt.Println("Hello ", name, "\n my name is ", customer.fullname)
}

func main() {
	agis := Customer{
		username: "agis",
		fullname: "Agisna Revaldo",
		age:      22,
	}
	fmt.Println(agis)

	budi := Customer{username: "budi", fullname: "budi santoso", age: 32}
	fmt.Println(budi)

	agis.sayhello(budi.fullname)
	budi.sayhello(agis.fullname)
}
