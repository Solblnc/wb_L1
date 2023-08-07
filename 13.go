package main

import "fmt"

// Поменять местами два числа без создания временной переменной

// JustChange - самый простой способ
func JustChange(a, b int) (int, int) {
	a, b = b, a
	return a, b
}

// MathChange - математический метод с вычитанием
func MathChange(a, b int) (int, int) {
	a += b
	b = a - b
	a = a - b

	return a, b
}

func main() {
	a := 7
	b := 93

	fmt.Println(JustChange(a, b))
	fmt.Println(MathChange(a, b))

}
