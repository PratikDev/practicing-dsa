// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/

package main

import "fmt"

func main() {
	arr := []int{5, 7, 7, 7, 8, 8, 10}
	target := 0

	first := getOccurance(arr, target, true)
	last := getOccurance(arr, target, false)

	fmt.Println("Ans:", []int{first, last})
}

func getOccurance(arr []int, target int, isFirstOccurance bool) int {
	ans := -1

	start := 0
	end := len(arr) - 1

	for start <= end {
		mid := (start + end) / 2

		if target < arr[mid] {
			end = mid - 1
		} else if target > arr[mid] {
			start = mid + 1
		} else {
			ans = mid

			if isFirstOccurance {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}
	}

	return ans
}
