package main

import "fmt"

// Car - родительская стуктура
type Car struct {
	Name string
	HP   int
}

// Ferrari - дочерняя структура
type Ferrari struct {
	Car
}

// Реализация наследования
func main() {
	Fr := Ferrari{Car{
		Name: "Ferrari",
		HP:   700,
	}}

	fmt.Println(Fr.Name, Fr.HP)
}
