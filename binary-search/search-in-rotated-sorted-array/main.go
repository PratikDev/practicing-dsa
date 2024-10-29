// https://leetcode.com/problems/search-in-rotated-sorted-array/description/

package main

import (
	"fmt"
)

func main() {
	arr := []int{4, 5, 6, 7, 0, 1, 2}
	target := 4
	result := search(arr, target)
	fmt.Println("Result:", result)
}

func search(arr []int, target int) int {
	pivot := findPivot(arr)

	ans := orderAgnosticBinarySearch(arr, target, 0, pivot)
	if ans == -1 {
		ans = orderAgnosticBinarySearch(arr, target, pivot+1, len(arr)-1)
	}

	return ans
}

func findPivot(arr []int) int {
	start := 0
	end := len(arr) - 1

	for start <= end {
		mid := (start + end) / 2

		// case 1: if mid is greater than the
		// next element, ans is mid
		if mid < end && arr[mid] > arr[mid+1] {
			return mid
		}

		// case 2: if mid is smaller than the
		// previous element, ans is the previous
		// element
		if mid > start && arr[mid] < arr[mid-1] {
			return mid - 1
		}

		if arr[start] >= arr[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return -1
}

func orderAgnosticBinarySearch(arr []int, target, start, end int) int {
	for start <= end {
		isAsc := arr[start] < arr[end]

		mid := (start + end) / 2

		if arr[mid] < target {
			if isAsc {
				start = mid + 1
			} else {
				end = mid - 1
			}
		} else if arr[mid] > target {
			if isAsc {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			return mid
		}
	}

	return -1
}
