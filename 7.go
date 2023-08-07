package main

import (
	"errors"
	"log"
	"sync"
)

// Реализовать конкурентную запись данных в map

type SafeMap struct {
	sync.RWMutex
	numbers map[int]int
}

func (s *SafeMap) Add(num int) {
	s.Lock()
	defer s.Unlock()
	s.numbers[num] = num
}

func (s *SafeMap) Get(num int) (int, error) {
	s.RLock()
	defer s.RUnlock()

	if number, ok := s.numbers[num]; ok {
		return number, nil
	}

	return 0, errors.New("Number does not exist")
}
func generateMap() {
	wg := sync.WaitGroup{}

	SafeMap := &SafeMap{
		numbers: map[int]int{},
	}

	// Пишем в мапу
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			SafeMap.Add(i)
		}(i)
	}

	// Читаем из мапы
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			num, err := SafeMap.Get(i)
			if err != nil {
				log.Println(err)
			} else {
				log.Println(num)
			}

		}(i)
	}
	wg.Wait()

}

func main() {
	generateMap()

}
