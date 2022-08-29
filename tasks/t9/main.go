package main

import (
	"fmt"
	"sync"
)

func main() {

	arr := []int{1, 2, 3, 4, 5}

	ch1 := make(chan int)
	ch2 := make(chan int)

	var wg sync.WaitGroup

	wg.Add(3)
	go func() {
		for _, v := range arr {
			ch1 <- v
		}
		close(ch1)
		wg.Done()
	}()

	go func() {
		for i := range ch1 {
			ch2 <- (i * i)
		}
		close(ch2)
		wg.Done()
	}()

	go func() {
		for i := range ch2 {
			fmt.Println(i)
		}
		wg.Done()
	}()
	wg.Wait()
}
