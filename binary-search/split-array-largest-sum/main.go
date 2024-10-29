// https://leetcode.com/problems/split-array-largest-sum/
package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{7, 2, 5, 10, 18}
	k := 3

	result := splitArray(arr, k)
	fmt.Println("Result:", result)
}

func splitArray(nums []int, k int) int {
	start := 0
	end := 0

	for _, num := range nums {
		start = int(math.Max(float64(start), float64(num)))
		end += num
	}

	for start < end {
		mid := (start + end) / 2

		sum := 0
		numOfSubArrays := 1

		for _, num := range nums {
			if (sum + num) > mid {
				sum = num
				numOfSubArrays++
			} else {
				sum += num
			}
		}

		if numOfSubArrays <= k {
			end = mid
		} else {
			start = mid + 1
		}
	}

	return start
}
