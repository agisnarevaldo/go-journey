package main

import "fmt"

func endApp() {
	fmt.Println("End app")
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("Shit Happen")
	}
	fmt.Println("App Running")
}

func main() {
	runApp(true)
}
