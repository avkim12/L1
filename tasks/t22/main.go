package main

import (
	"fmt"
	"math/big"
)

func main() {

	a := new(big.Int)
	b := new(big.Int)
	c := new(big.Int)

	a.SetString("9223372036854775808", 10)
	b.SetString("9223372036854775809", 10)

	fmt.Println(c.Mul(a, b))
	fmt.Println(c.Div(a, b))
	fmt.Println(c.Add(a, b))
	fmt.Println(c.Sub(a, b))
}