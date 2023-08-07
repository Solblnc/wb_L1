package main

import (
	"fmt"
	"strconv"
)

// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0

func IntToBits(num int) string {
	return strconv.FormatInt(int64(num), 2)
}

func setBit(n int, pos int) int {
	pos = len(IntToBits(n)) - pos
	mask := (1 << pos)
	n = n | mask
	return n
}

func clearBit(n int, pos int) int {
	pos = len(IntToBits(n)) - pos
	mask := ^(1 << pos)
	n &= mask
	return n
}

func main() {
	fmt.Println(clearBit(981, 2)) // 725
	fmt.Println(setBit(273, 2))   // 401
}
