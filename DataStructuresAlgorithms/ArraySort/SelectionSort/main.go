package main

import "fmt"

/*
 * Implement Selection Sort
 */

func SelectionSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minLoc := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minLoc] {
				minLoc = j
			}
		}
		if minLoc != i {
			arr[i], arr[minLoc] = arr[minLoc], arr[i]
		}
	}
	return arr
}

func main() {
	testArray := []int{1, 5, 3, 7, 6, 9, 8, 2, 4}

	fmt.Println(SelectionSort(testArray))
}
