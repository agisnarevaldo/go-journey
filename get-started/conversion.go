package main

import "fmt"

func main() {
	const nilai32 int32 = 32456
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai64)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	var name = "Agisna Revaldo"
	var n = name[0]
	var nString = string(n)

	fmt.Println(name)
	fmt.Println(n)
	fmt.Println(nString)
}
