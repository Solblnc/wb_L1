package main

import (
	"fmt"
	"strings"
)

//Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow».

func reverseWords(s string) string {
	// Разделяем строка на слова (почти как strings.Split но лучше))))
	words := strings.Fields(s)
	// Два указателя
	left, right := 0, len(words)-1

	for left < right {
		// Меняем местами
		words[left], words[right] = words[right], words[left]
		// ++ / --
		left++
		right--
	}

	// Соединияем в строку все наши слова
	return strings.Join(words, " ")
}

func main() {
	s := "Hit me up"
	fmt.Println(reverseWords(s))

}
