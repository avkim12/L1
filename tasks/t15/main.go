package main

import (
	"fmt"
	"strings"
)

// в строке могут быть символы не входящие в таблицу ASCII,
// тогда в justString запишутся неправильные символы

// var justString string
// func someFunc() {
//   v := createHugeString(1 << 10)
//   justString = v[:100]
// }

// func main() {
//   someFunc()
// }

func createHugeString(i int) string {
	return strings.Repeat("A", i)
}

func someFunc() {
	v := createHugeString(1 << 10)
	runes := []rune(v)
	justString := string(runes[:100])
	fmt.Println(justString)
}

func main() {
	someFunc()
}
