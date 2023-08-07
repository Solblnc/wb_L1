package main

import "fmt"

//Разработать программу, которая переворачивает подаваемую на ход строку
//(например: «главрыба — абырвалг»). Символы могут быть unicode.

func reverseLettersFirst(s string) string {
	// Переводим в руны
	arr := []rune(s)
	// Два указателя - левый и правый
	left, right := 0, len(arr)-1

	for left < right {
		// Меняем местами
		arr[left], arr[right] = arr[right], arr[left]
		// Левый инкрементируем, правый дикрементируем
		left++
		right--
	}

	return string(arr)
}

func main() {
	s := "magazine"
	s1 := "привет"
	fmt.Println(reverseLettersFirst(s))
	fmt.Println(reverseLettersFirst(s1))
}
