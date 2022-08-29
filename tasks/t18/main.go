package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	m sync.Mutex
	i int
}

func (c *Counter) Increment() {
	c.m.Lock()
	defer c.m.Unlock()
	c.i++
}

func (c *Counter) Get() int {
	c.m.Lock()
	defer c.m.Unlock()
	return c.i
}

func main() {

	c := Counter{}

	var wg sync.WaitGroup

	for i := 1; i <= 123; i++ {
		wg.Add(1)
		go func() {
			c.Increment()
			wg.Done()
		}()

	}
	wg.Wait()
	
	fmt.Println(c.Get())
}