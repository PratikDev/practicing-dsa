package main

import "fmt"

func main() {
	arr := []int{2, 4, 6, 9, 10, 16, 27, 33}
	target := 5

	ceiling := findCeiling(arr, target)
	fmt.Println("Ceiling:", ceiling)

	floor := findFloor(arr, target)
	fmt.Println("Floor:", floor)
}

func findCeiling(arr []int, target int) int {
	start := 0
	end := len(arr) - 1

	// EDGE CASE: if target is greater than
	// even the last element of the slice
	// return -1
	if target > arr[end] {
		return -1
	}

	// basic binary search
	for start <= end {
		mid := (start + end) / 2

		if target == arr[mid] {
			return arr[mid]
		}

		if target < arr[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	// only difference is we're not returning -1
	// when not found. instead, we're returning the
	// upper bound (start)
	return arr[start]
}

func findFloor(arr []int, target int) int {
	start := 0
	end := len(arr) - 1

	// EDGE CASE: if target is less than
	// even the first element of the slice
	// return -1
	if target < arr[0] {
		return -1
	}

	// basic binary search
	for start <= end {
		mid := (start + end) / 2

		if target == arr[mid] {
			return arr[mid]
		}

		if target < arr[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	// only difference is we're not returning -1
	// when not found. instead, we're returning the
	// lower bound (end)
	return arr[end]
}
