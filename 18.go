package main

import (
	"fmt"
	"sync"
)

//Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде
//По завершению программа должна выводить итоговое значение счетчика.

type counter struct {
	count int
	sync.Mutex
}

func (c *counter) Increment() {
	c.Lock()
	defer c.Unlock()
	c.count++
}

func main() {
	wg := sync.WaitGroup{}
	counter := counter{count: 0}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			counter.Increment()
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(counter.count)

}
