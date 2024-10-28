package main

import "fmt"

func main() {
	arr := []int{2, 3, 6, 8, 10, 12, 13, 19, 22, 25, 28, 31, 33, 37, 40, 46, 55, 75, 81}
	target := 9

	start := 0
	end := 1

	// if target is greater than end value
	// means we need to change the chunk
	for target > arr[end] {
		temp := start
		start = end + 1
		end = end + (end-temp+1)*2
	}

	index := binarySearch(arr, target, start, end)
	fmt.Println("Index:", index)
}

// basic binary search
func binarySearch(arr []int, target, start, end int) int {
	for start <= end {
		midIndex := (start + end) / 2

		if target < arr[midIndex] {
			end = midIndex - 1
		} else if target > arr[midIndex] {
			start = midIndex + 1
		} else {
			return midIndex
		}
	}

	return -1
}
