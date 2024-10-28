// https://leetcode.com/problems/peak-index-in-a-mountain-array/description/

package main

import "fmt"

func main() {
	arr := []int{2, 3, 5, 9, 7, 0}

	result := peakIndexInMountainArray(arr)
	fmt.Println("Result:", result)
}

func peakIndexInMountainArray(arr []int) int {
	start := 0
	end := len(arr) - 1

	for start <= end {
		mid := (start + end) / 2

		leftSmall := arr[mid] > arr[mid-1]
		rightSmall := arr[mid] > arr[mid+1]

		// if mid is the peak
		if leftSmall && rightSmall {
			return arr[mid]
		}

		if leftSmall {
			start = mid + 1
		} else if rightSmall {
			end = mid
		}
	}

	return arr[end]
}
