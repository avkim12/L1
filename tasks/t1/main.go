package main

import "fmt"

type Human struct {
}

func (h *Human) Greet() {
	fmt.Println("Hello!")
}

type Action struct {
	Human
}

func main() {
	var a Action

	a.Greet()
}
