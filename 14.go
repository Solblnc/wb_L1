package main

import (
	"fmt"
)

// Разработать программу, которая в рантайме способна определить тип переменной:
//int, string, bool, channel из переменной типа interface{}.

func GetType(value interface{}) string {
	switch value.(type) {
	case int:
		return fmt.Sprintf("%T\n", value)
	case string:
		return fmt.Sprintf("%T\n", value)
	case bool:
		return fmt.Sprintf("%T\n", value)
	case chan string:
		return fmt.Sprintf("%T\n", value)
	default:
		return fmt.Sprintf("Unknown type: %T\n", value)

	}

}

func main() {
	number := 5
	letter := "5"
	ch := make(chan string)
	fmt.Println(GetType(number))
	fmt.Println(GetType(letter))
	fmt.Println(GetType(ch))
}
