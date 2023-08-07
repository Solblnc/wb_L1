package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}

	result := 0

	for i := range nums {
		wg.Add(1)

		go func(x int) {
			defer wg.Done()
			result += x * x
		}(nums[i])
	}
	wg.Wait()

	fmt.Println(result)
}
