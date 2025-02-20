package main

import (
	"fmt"
)

type NoKTP string

func main() {
	var template NoKTP = "xxxxxxxx"

	var ktp1 NoKTP = NoKTP(template)

	fmt.Println(ktp1)
	fmt.Printf("Type: %T Value: %v", ktp1, ktp1)

}
