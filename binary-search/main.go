package main

import "fmt"

func main() {
	arr := []int{5, 7, 7, 7, 7, 8, 8, 10}
	target := 7

	index := binarySearchWithRecursion(arr, target, 0, len(arr)-1)
	fmt.Println("Index With Recursion:", index)

	_index := binarySearchWithLoop(arr, target)
	fmt.Println("Index With Loop:", _index)
}

func binarySearchWithRecursion(arr []int, target, start, end int) int {
	if start > end {
		return -1
	}

	mid := (start + end) / 2

	if arr[mid] == target {
		return mid
	} else if target < arr[mid] {
		return binarySearchWithRecursion(arr, target, start, mid-1)
	} else {
		return binarySearchWithRecursion(arr, target, mid+1, end)
	}
}

func binarySearchWithLoop(arr []int, target int) int {
	start := 0
	end := len(arr) - 1

	for start <= end {
		midIndex := (start + end) / 2
		mid := arr[midIndex]

		if target < mid {
			end = midIndex - 1
		} else if target > mid {
			start = midIndex + 1
		} else {
			return midIndex
		}
	}

	return -1
}
