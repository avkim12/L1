package main

import "fmt"

func main() {
	a1 := []string{"cat", "cat", "dog", "cat", "tree"}
	a2 := []string{"koala", "horse", "dog", "cat", "tree"}
	
	set := make(map[string]struct{})
	res := make([]string, 0)

	for _, v := range a1 {
		set[v] = struct{}{}
	}

	for _, v := range a2 {
		if _, ok := set[v]; ok {
			res = append(res, v)
		}
	}

	fmt.Println(res)
}