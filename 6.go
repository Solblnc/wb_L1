package main

// Реализовать все возможные способы остановки выполнения горутины

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	closeCh()
	signal()
	withContext()
}

// Просто закрыв канал можно завершить ее работу
func closeCh() {

	wg := sync.WaitGroup{}

	ch := make(chan int)

	go func() {

		wg.Add(1)
		defer wg.Done()

		for {
			_, ok := <-ch
			if !ok {
				fmt.Println("Goroutine closeCh done")
				return
			}
		}
	}()
	ch <- 1
	ch <- 1
	close(ch)
	wg.Wait()

}

// Горутина завершит свою работу при получении сигнала из канала
func signal() {
	ch := make(chan bool)
	wg := sync.WaitGroup{}

	go func() {
		wg.Add(1)
		defer wg.Done()

		for {
			select {
			case <-ch:
				fmt.Println("Goroutine signal done")
				return
			}
		}
	}()
	ch <- true
	wg.Wait()

}

// Горутину можно завершить закрытием контекста через cancel функцию
func withContext() {
	ch := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				ch <- 0
				return

			}
		}
	}(ctx)
	cancel()

	<-ch
	fmt.Println("Goroutine withCancel done")
}
