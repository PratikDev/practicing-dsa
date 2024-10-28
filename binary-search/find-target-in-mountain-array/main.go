// https://leetcode.com/problems/find-in-mountain-array/description/

package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 3, 1}
	target := 2

	ans := findInMountainArray(target, arr)
	fmt.Println("Ans:", ans)
}

func findInMountainArray(target int, mountainArr []int) int {
	peak := peakIndexInMountainArray(mountainArr)

	// do binary search in LHS of peak (inclusive)
	ans := orderAgnosticBinarySearch(mountainArr, target, 0, peak)
	if ans == -1 {
		// if not found in last step
		// do binary search in RHS of peak (exclusive)
		ans = orderAgnosticBinarySearch(mountainArr, target, peak+1, len(mountainArr)-1)
	}

	return ans
}

func peakIndexInMountainArray(arr []int) int {
	start := 0
	end := len(arr) - 1

	for start <= end {
		mid := (start + end) / 2

		leftSmall := arr[mid] > arr[mid-1]
		rightSmall := arr[mid] > arr[mid+1]

		if leftSmall && rightSmall {
			return mid
		}

		if leftSmall {
			start = mid
		} else if rightSmall {
			end = mid
		}
	}

	return end
}

func orderAgnosticBinarySearch(arr []int, target, start, end int) int {
	isAsc := arr[start] < arr[end]

	for start <= end {
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
