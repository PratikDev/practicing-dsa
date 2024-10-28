// https://leetcode.com/problems/peak-index-in-a-mountain-array/description/

package main

import "fmt"

func main() {
	arr := []int{2, 3, 5, 9, 7, 1}

	result := peakIndexInMountainArray(arr)
	fmt.Println("Result:", result)
}

func peakIndexInMountainArray(arr []int) int {
	start := 0
	end := len(arr) - 1

	for start <= end {
		mid := (start + end) / 2

		// if mid is the peak
		if (arr[mid] > arr[mid-1]) && (arr[mid] > arr[mid+1]) {
			return arr[mid]
		}

		if arr[mid] > arr[mid-1] {
			start = mid
		} else if arr[mid] > arr[mid+1] {
			end = mid
		}
	}

	return arr[end]
}
