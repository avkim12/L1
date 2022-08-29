package main

import "fmt"

func binarySearch(needle int, haystack []int) int {

	low := 0
	high := len(haystack) - 1

	for low <= high {
		median := (low + high) / 2

		if haystack[median] < needle {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(haystack) || haystack[low] != needle {
		return -1
	}

	return low
}

func main() {

	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(binarySearch(2, arr))
}
