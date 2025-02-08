package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a number between 1 and 10: ")
	targetStr, _ := reader.ReadString('\n')
	targetStr = strings.TrimSpace(targetStr)
	target, err := strconv.Atoi(targetStr)
	if target > len(arr)+1 {
		fmt.Println("Invalid input, please enter a number between 1 and 10")
		return
	}
	if err != nil {
		fmt.Println("Invalid input, please enter a number")
		return
	}
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
