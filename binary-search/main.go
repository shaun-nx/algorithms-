package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 6
	fmt.Println(myExample(arr, target))
}

func myExample(arr []int, target int) int {
	// Define boundary of the  list
	low, high := 0, len(arr)-1
	// Loop until the low index is less than or equal to the high index
	for low <= high {
		// Calculate the middle index
		mid := low + (high-low)/2
		// If the middle element is equal to the target, return the middle index
		if arr[mid] == target {
			return mid
		}
		// If the middle element is less than the target, set the low index to mid + 1
		if arr[mid] < target {
			low = mid + 1
		} else {
			// If the middle element is greater than the target, set the high index to mid - 1
			high = mid - 1
		}
	}
	return -1
}
