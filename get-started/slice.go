package main

import "fmt"

func main() {
	months := [12]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}

	var slice1 []string = months[4:7]
	fmt.Println(slice1)

	slice2 := months[:3]
	fmt.Println(slice2)

	slice3 := months[10:]
	fmt.Println(slice3)

	slice4 := months[:]
	fmt.Println(slice4)

	slice2[0] = "Januari"
	slice2[1] = "Februari"
	fmt.Println(slice2)
	fmt.Println(months)

	slice5 := append(slice4, "Happy New Year")
	fmt.Println(slice5)
	fmt.Println(slice4)

	var newSlice []string = make([]string, 3, 5)
	newSlice[0] = "Muhamad"
	newSlice[1] = "Agisna"
	newSlice[2] = "Revaldo"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Paleee")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))
	newSlice2 = append(newSlice2, "Brian")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	toSlice := make([]string, len(newSlice2), cap(newSlice2))
	copy(toSlice, newSlice2)

	fmt.Println(toSlice)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	iniArray := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(iniArray)
	fmt.Println(len(iniArray))
	fmt.Println(cap(iniArray))

	iniSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(iniSlice)
	fmt.Println(len(iniSlice))
	fmt.Println(cap(iniSlice))
}
