package main

import (
	"fmt"
	"sync"
)

func main() {
	var sum int
	arr := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup
	wg.Add(len(arr))

	for _, v := range arr {
		go func(v int) {
			sum += v * v
			wg.Done()
		}(v)
	}
	wg.Wait()
	fmt.Println(sum)
}
