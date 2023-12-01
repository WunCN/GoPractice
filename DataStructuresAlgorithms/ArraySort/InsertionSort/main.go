package main

import "fmt"

/*
 * Implement Insertion Sort
 */

func InsertionSort(arr []int) []int {
	n := len(arr)
	for i := 1; i < n; i++ {
		tmp := arr[i]
		j := i
		for j > 0 && arr[j-1] > tmp {
			arr[j] = arr[j-1]
			j--
		}
		arr[j] = tmp
	}
	return arr
}

func main() {
	testArray := []int{1, 5, 3, 7, 6, 9, 8, 2, 4}

	fmt.Println(InsertionSort(testArray))
}
