package main

import "fmt"

func main() {
	a, b := 0, 1
	a, b = b, a
	fmt.Println(a, b)
}
