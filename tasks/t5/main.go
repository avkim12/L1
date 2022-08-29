package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	c := make(chan int, 1)

	var n int
	fmt.Scan(&n)

	time.AfterFunc(time.Duration(n)*time.Second, func() {
		os.Exit(0)
	})

	for {
		c <- rand.Int()
		fmt.Println(<-c)
	}
}
