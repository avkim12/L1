package main

import (
	"fmt"
	"math/big"
)

func main() {

	a := new(big.Int)
	b := new(big.Int)
	c := new(big.Int)

	a.SetString("98765432123456789", 10)
	b.SetString("12345678987654321", 10)

	fmt.Println(c.Mul(a, b))
	fmt.Println(c.Div(a, b))
	fmt.Println(c.Add(a, b))
	fmt.Println(c.Sub(a, b))
}