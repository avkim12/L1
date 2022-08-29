package main

import (
	"fmt"
	"strings"
)

func main() {

	var str string
	fmt.Scan(&str)

	str = strings.ToLower(str)

	m := make(map[rune]struct{})
	for _, v := range str {
		_, exist := m[v]
		if exist {
			fmt.Println(false)
			return
		}
		m[v] = struct{}{}
	}
	fmt.Println(true)
}
