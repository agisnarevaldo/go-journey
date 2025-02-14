package main

import "fmt"

var i, j int = 2, 4

func main() {
	var (
		c      = true
		python = false
		java   = "not"
	)

	fmt.Println(c, python, java, i, j)
}
