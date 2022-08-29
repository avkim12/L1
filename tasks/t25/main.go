package main

import "time"

func sleepBefore(dur time.Duration) {
	finishAt := time.Now().Add(dur)

	for time.Now().Before(finishAt) {
	}
}

func wakeUpAfter(dur time.Duration) {
	<-time.After(dur)
}

func main() {
	sleepBefore(time.Second * time.Duration(1))
	wakeUpAfter(time.Second * time.Duration(2))
}
