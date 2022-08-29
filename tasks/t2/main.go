package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup
	wg.Add(len(arr))

	for _, v := range arr {
		go func(v int) {
			fmt.Println(v * v)
			wg.Done()
		}(v)
	}
	wg.Wait()
}
