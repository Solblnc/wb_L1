package main

import "fmt"

// Реализовать пересечение двух неупорядоченных множеств

// Множество почти то же самое, что и мапа но без сохранения значения
// Удобно использовать чтобы быстро найти вхождение элемента

// CreateSet -  создает множестов множество
func CreateSet(array []int) map[int]struct{} {
	set := make(map[int]struct{})

	for _, value := range array {
		set[value] = struct{}{}
	}

	return set
}

// Intersection - ищет пересечение в двух множествах
func Intersection(set1, set2 map[int]struct{}) map[int]struct{} {
	result := make(map[int]struct{})

	for key, _ := range set1 {
		if _, ok := set2[key]; ok {
			result[key] = struct{}{}
		}
	}

	return result
}

func main() {
	array1 := []int{1, 2, 3, 4, 5, 14, 38, 4}
	array2 := []int{1, 2, 38, 4}
	set1 := CreateSet(array1)
	set2 := CreateSet(array2)
	fmt.Println(Intersection(set1, set2))
}
