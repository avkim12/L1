package main

import "fmt"

func main() {

	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(slice)

	var i int
	fmt.Scan(&i)

	slice = append(slice[:i], slice[i+1:]...)
	fmt.Println(slice)
}
