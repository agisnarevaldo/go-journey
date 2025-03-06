package main

import "fmt"

func main() {
	//var person map[string]string = map[string]string{}
	//person["name"] = "Agisna Revaldo"
	//person["address"] = "Tasikmalaya"

	person := map[string]string{
		"name":    "Agisna",
		"address": "Tasikmalaya",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	book := make(map[string]string)
	book["title"] = "Belajar GoLang"
	book["author"] = "Agisna Revaldo"
	book["Ups"] = "Salah"

	fmt.Println(book)
	delete(book, "Ups")
	fmt.Println(book)
}
