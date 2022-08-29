package main

import (
	"context"
	"fmt"
	"time"
)

func quitChan() {

	quit := make(chan bool)

	go func() {
		for {
			select {
			case <-quit:
				return
			default:
				fmt.Println("GO")
			}
		}
	}()
	time.Sleep(time.Second)
	quit <- true
}

func contextWithCancel() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				fmt.Println("GO")
			}
		}
	}(ctx)
	time.Sleep(time.Second)
}

func contextWithTimeout() {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				fmt.Println("GO")
			}
		}
	}(ctx)
	time.Sleep(time.Second)
}

func main() {
	quitChan()
	contextWithCancel()
	contextWithTimeout()
}
