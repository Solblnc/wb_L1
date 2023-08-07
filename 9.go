package main

// Разработать конвейер чисел
// Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2,
// После чего данные из второго канала должны выводиться в stdout

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	ch1 := make(chan int)
	ch2 := make(chan int)
	wg := sync.WaitGroup{}

	// Пишем во второй канал квадраты чисел
	go func() {
		wg.Add(1)
		defer wg.Done()

		for num := range ch1 {
			ch2 <- num * num
		}
		close(ch2)

	}()

	// Читаем из второго канала значения квадратов чисел
	go func() {
		wg.Add(1)
		defer wg.Done()
		for num := range ch2 {
			fmt.Printf("%d ", num)
		}
	}()

	// Пишем значения в канал с обычными числами
	for _, value := range nums {
		ch1 <- value
	}
	close(ch1)
	wg.Wait()

}
