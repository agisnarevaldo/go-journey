package main

import "fmt"

func getGoodBye(name string) string {
	return "Bye " + name
}

func main() {
	sayBye := getGoodBye("Agis")
	fmt.Println(sayBye)

	seeYou := getGoodBye
	fmt.Println(seeYou("Revaldo"))
}
