package main

import (
	"fmt"
	"math"
	"math/big"
)

// Разработать программу, которая перемножает, делит, складывает, вычитает две числовых
// переменных a, b, значение которых > 2^20.

func main() {
	num1 := math.Pow(2, 45)
	num2 := math.Pow(2, 22)
	a := big.NewInt(int64(num1))
	b := big.NewInt(int64(num2))

	res := big.NewInt(0)

	fmt.Println("a*b = ", res.Mul(a, b))
	fmt.Println("a/b = ", res.Div(a, b))
	fmt.Println("a-b = ", res.Sub(a, b))
	fmt.Println("a+b = ", res.Add(a, b))
}
