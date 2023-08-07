package main

import "fmt"

// Удалить i-ый элемент из слайса.

func deleteElement(nums []int, target int) []int {
	// Если слайс меньше двух - просто возвращаем его
	if len(nums) < 1 {
		return nums
	}

	if len(nums) < target {
		return []int{}
	}
	// Возвращаем слайс до элемента, который нужно удалить
	// И возращаем часть слайса после этого элемента
	// Получаем слайс с удаленным элементом
	return append(nums[:target], nums[target+1:]...)

}

func main() {
	nums := []int{1, 3, 5, 7, 9}
	target := 3
	fmt.Println(deleteElement(nums, target))

}
