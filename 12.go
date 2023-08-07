package main

import "fmt"

// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество

// CreateSetString - создает множество из массива
func CreateSetString(array []string) map[string]struct{} {
	set := make(map[string]struct{})

	for _, value := range array {
		set[value] = struct{}{}
	}

	return set

}

func main() {
	array := []string{"cat, cat, dog, cat, tree"}
	fmt.Println(CreateSetString(array))

}
