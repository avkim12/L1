package main

import "fmt"

func main() {
	var a, i, v int

	fmt.Scan(&i)
	i--
	a = 123
	fmt.Printf("%b\n", a)

	v = (a >> i) & 1 // i-й бит числа
	if v == 1 {
		a ^= 1 << i
		fmt.Printf("%b\n", a)
	} else if v == 0 {
		a |= 1 << i
		fmt.Printf("%b\n", a)
	}
}
