package main

import "fmt"

func endApp() {
	fmt.Println("End App.")
	message := recover()
	fmt.Println("Terjadi panic: ", message)
}

func runApp(error bool) {
	defer endApp()
	fmt.Println("App Running")
	if error {
		panic("ERROR")
	}
}

func main() {
	runApp(false)
}
