package main

import "fmt"

//Реализовать бинарный поиск встроенными методами языка.

func binarySearch(array []int, target int) int {
	// Создаем два указателя (левый и правый)
	left, right := 0, len(array)-1
	// Левый должен быть меньше или равен правому
	for left <= right {
		// Создаем среднее значение (середина массива - массив отсортированный)
		mid := (left + right) / 2
		// Сравниваем значения с целью, передвигаем указатели
		if array[mid] < target {
			left = mid + 1
		} else if array[mid] > target {
			right = mid - 1
		} else {
			return mid
		}
	}

	// Возвращаем -1 если такого элемента не существует
	return -1
}

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7}
	target := 4
	fmt.Println(binarySearch(array, target))

}
