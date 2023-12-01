package main

import "fmt"

/*
 * Implement Bubble Sort
 */

func BubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		changed := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				changed = true
			}
		}
		if !changed {
			break // pruning: not changed means the array is already ordered
		}
	}
	return arr
}

func BubbleSortLastSwap(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		lastSwapped := -1
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				lastSwapped = j
			}
		}
		if lastSwapped == -1 {
			break
		}
		n = lastSwapped + 1 // array after lastSwapped is ordered, just stop at that location
	}
	return arr
}

func main() {
	testArray := []int{1, 5, 3, 7, 6, 9, 8, 2, 4}

	fmt.Println(BubbleSort(testArray))
	fmt.Println(BubbleSortLastSwap(testArray))
}
