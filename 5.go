package main

// Разработать программу, которая будет последовательно отправлять значения в канал,
// А с другой стороны канала — читать
// По истечению N секунд программа должна завершаться
import (
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	var n int
	fmt.Println("Enter amount of time: ")
	fmt.Scan(&n)

	ch := make(chan int, n)
	wg := sync.WaitGroup{}

	// Запуск воркеров

	go func() {
		defer wg.Done()
		for v := range ch {
			fmt.Println(v)
		}

	}()

	// Создаем канала для сигнала прерывания (таймер)
	timer := time.After(time.Duration(n) * time.Second)
	i := 0

	// В отдельной горутине ждем истечения таймера
	go func() {
		<-timer
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
