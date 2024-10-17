package main

import (
	"fmt"
)

// merge merges two sorted slices into a single sorted slice
func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	// Merge the two slices into result
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Append any remaining elements from left or right
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

// mergeSort recursively sorts the slice using merge sort
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// Split the array into two halves
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	// Merge the sorted halves
	return merge(left, right)
}

func main() {
	// Example array to be sorted
	arr := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Println("Original array:", arr)

	// Perform merge sort
	sortedArr := mergeSort(arr)
	fmt.Println("Sorted array:", sortedArr)
}
