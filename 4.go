package main

// Реализовать постоянную запись данных в канал (главный поток).
// Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
// Необходима возможность выбора количества воркеров при старте.

// Программа должна завершаться по нажатию Ctrl+C.
// Выбрать и обосновать способ завершения работы всех воркеров.
import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	var n int
	fmt.Println("Enter amount of workers: ")
	fmt.Scan(&n)

	ch := make(chan int, n)
	wg := sync.WaitGroup{}

	// Запуск воркеров
	for i := 0; i < n; i++ {

		go func(index int) {
			defer wg.Done()
			for v := range ch {
				fmt.Printf("Worker: %d: %v\n", index, v)
			}

		}(i)
	}

	// Создаем канала для сигнала прерывания
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	i := 0

	// В отдельной горутине ждем сигнала
	go func() {
		<-c
		close(ch)
		wg.Wait()
		os.Exit(1)
	}()

	// Отправляем данные в канал
	for {
		ch <- i
		i++
		time.Sleep(200 * time.Millisecond)
	}

}
