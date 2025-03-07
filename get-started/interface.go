package main

import "fmt"

// INTERFACE:
// Definisi
type Creatures interface {
	Name() string
	Sound() string
}

// Implementasi
type Dog struct {
	name string
}

func (d Dog) Name() string {
	return d.name
}

func (d Dog) Sound() string {
	return "Wooff!"
}

type Human struct {
	name string
}

func (h Human) Name() string {
	return h.name
}

func (h Human) Sound() string {
	return "Hallo saya " + h.name
}

func main() {
	// Penggunaan
	home := []Creatures{
		Human{name: "Andy"},
		Dog{name: "Buggy"},
	}

	for _, creatures := range home {
		fmt.Printf("%v: %v\n", creatures.Name(), creatures.Sound())
	}
}
