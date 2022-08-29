package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
)

func worker(c chan int) {

	for {
		a := <-c
		fmt.Println(a)
	}
}

func main() {

	c := make(chan int)
	sc := make(chan os.Signal, 1)

	var n int
	fmt.Scan(&n)

	signal.Notify(sc, syscall.SIGINT)

	for i := 0; i < n; i++ {
		go worker(c)
	}

	for {
		select {
		case <-sc:
			os.Exit(0)
		default:
			c <- rand.Int()
		}
	}
}
