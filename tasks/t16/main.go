package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	left, right := 0, len(arr)-1

	// Pick a pivot
	pivotIndex := (left + right) / 2

	// Move the pivot to the right
	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]

	// Pile elements smaller than the pivot on the left
	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	// Place the pivot after the last smaller element
	arr[left], arr[right] = arr[right], arr[left]

	// Go down the rabbit hole
	quickSort(arr[:left])
	quickSort(arr[left+1:])

	return arr
}

func main() {

	arr := []int{-2, 0, 4, 6, -7, 3, 2, -5, 7, 3}

	fmt.Println(quickSort(arr))
}
