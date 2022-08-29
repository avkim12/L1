package main

import "fmt"

type Animal interface {
	SayMeow()
}

type Cat struct {
}

func (c Cat) SayMeow() {
	fmt.Println("Meow")
}

type Dog struct {
	Cat
}

func (d *Dog) SayWoof() {
	fmt.Println("Woof")
}

type AnimalTrainer struct {
}

func (at *AnimalTrainer) ConvertToMeow(a Animal) {
	a.SayMeow()
}

type DogAdapter struct {
	d *Dog
}

func (da DogAdapter) SayMeow() {
	da.d.SayMeow()
}

func main() {

	at := &AnimalTrainer{}
	d := &Dog{}
	da := &DogAdapter{d: d}
	at.ConvertToMeow(da)
}
