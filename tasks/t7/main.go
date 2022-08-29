package main

import (
	"fmt"

	"sync"
	"time"
)

func set(m *sync.Mutex, dst map[int]string, v string) {
	
	for k := 0; k < 5; k++ {
		m.Lock()
		dst[k] = v
		m.Unlock()
		time.Sleep(time.Millisecond * 200)
	}
}

func main() {

	var m sync.Mutex

	dst := make(map[int]string)

	go set(&m, dst, "1")
	go set(&m, dst, "2")
	go set(&m, dst, "3")
	go set(&m, dst, "4")
	go set(&m, dst, "5")

	time.Sleep(time.Second)

	for i, v := range dst {
		fmt.Printf("key: %d, value: %s\n", i, v)
	}
}
