// https://leetcode.com/problems/find-smallest-letter-greater-than-target/description/

package main

import "fmt"

func main() {
	letters := []string{"c", "f", "j"}
	target := "a"

	result := nextGreatestLetter(letters, target)
	fmt.Println("Result:", result)
}

func nextGreatestLetter(letters []string, target string) string {
	start := 0
	end := len(letters) - 1

	// for cover the wraps
	if target >= letters[end] {
		return letters[0]
	}

	// basic binary search
	for start <= end {
		mid := (start + end) / 2

		if target < letters[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return letters[start]
}
