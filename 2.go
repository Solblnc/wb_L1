package main

import (
	"io"
	"os"
	"strconv"
	"sync"
)

func main() {
	nums := []int{2, 4, 6, 8, 10}
	// Создаем WaitGroup
	wg := sync.WaitGroup{}

	for i := range nums {
		// Добавляем счетчик для вейт группы
		wg.Add(1)
		go func(x int) {
			defer wg.Done() // Горутина выполнена

			io.WriteString(os.Stdout, strconv.Itoa(x*x)+"\n")
		}(nums[i])
	}
	// Ждем выполнения всех горутин
	wg.Wait()
}
