package main

import (
	"fmt"
	"strings"
)

//К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.
//
//
//var justString string
//func someFunc() {
//	v := createHugeString(1 << 10)
//	justString = v[:100]
//}
//
//func main() {
//	someFunc()
//}

// Создается строка из 2^10 символов которая потом обрезается до 100
// Самый простой способо - просто создать строку из 100 сиволов
// + использование глобальной переменной не лучшее решение т.к. неявное изменение состояний + установление зависимости,
// а значит неудобства при тестировании

func createHugeString(i int) string {
	return strings.Repeat("a", i)
}

func someFunc() string {
	v := createHugeString(1 << 10)
	// Переводим в руны для разделения по символам, а не по байтам
	runes := []rune(v)
	return string(runes[:100])
}

func main() {
	justString := someFunc()
	fmt.Println(justString, len([]rune(justString)))
}
