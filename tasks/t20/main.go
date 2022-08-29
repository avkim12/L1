package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "snow dog sun"
	fmt.Println(str)

	strs := strings.Split(str, " ")
	i := 0
	j := len(strs) - 1

	for i < j {
		strs[i], strs[j] = strs[j], strs[i]
		i++
		j--
	}

	newStr := strings.Join(strs, " ")
	fmt.Println(newStr)
}
