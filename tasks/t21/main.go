package main

import "fmt"

type Feline interface {
	SayMeow()
}

type Cat struct {
}

func (c Cat) SayMeow() {
	fmt.Println("Meow")
}

type Dog struct {
}

func (d Dog) ConvertToMeow() {
	fmt.Println("Meow")
}

type DogAdapter struct {
	d Dog
}

func (da DogAdapter) SayMeow() {
	da.d.ConvertToMeow()
}

func main() {

	da := &DogAdapter{}
	da.SayMeow()
}
