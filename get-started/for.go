package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("Perulangan ke: %v\n", i)
	}
	fmt.Println("Selesai")

	//for counter := 1; counter <= 10; {
	//	fmt.Println("terhitung: ", counter)
	//	counter++
	//}

	names := [...]string{"Muhamad", "Agisna", "Revaldo"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, value := range names {
		fmt.Println("Index:", index, "| name:", value)
	}

	for _, value := range names { // jika tidak butuh index bisa digante dengan _ (blank identifier/underscore)
		fmt.Println(value)
	}
}
