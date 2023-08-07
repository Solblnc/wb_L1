package main

import (
	"fmt"
	"math/rand"
)

//Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

func quickSort(array []int) []int {
	// Если длина массива меньше двух, нам нечего сортировать - возвращаем массив
	if len(array) < 2 {
		return array
	}
	// Создаем два указателя - начало и конец (левый и правый)
	left, right := 0, len(array)-1
	// Создаем пивот
	pivot := rand.Int() % len(array)
	// Передвигаем пивот к правому значению
	array[pivot], array[right] = array[right], array[pivot]
	// Передвигаем значения меньшие чем пивот налево
	for i := range array {
		if array[i] < array[right] {
			array[left], array[i] = array[i], array[left]
			left++
		}
	}
	// Ставим пивот сразу после последнего минимального значения
	array[left], array[right] = array[right], array[left]
	// Рекурсивно пробегаемся по всему массиву
	quickSort(array[:left])
	quickSort(array[left+1:])

	return array

}

func main() {
	array := []int{4, 8, 2, 6, 8, 75, 98, 34, 193}
	fmt.Println(quickSort(array))

}
