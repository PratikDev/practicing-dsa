// https://leetcode.com/problems/find-all-duplicates-in-an-array/description/

package main

import "fmt"

func main() {
	arr := []int{4, 3, 2, 7, 8, 2, 3, 1}
	result := findDuplicates(arr)
	fmt.Println(result)
}

func findDuplicates(nums []int) []int {
	helperArr := make([]int, len(nums)+1)

	for _, val := range nums {
		if helperArr[val] == 0 {
			helperArr[val] = 1
		} else {
			helperArr[val] = 2
		}
	}

	ret := []int{}
	for idx, val := range helperArr {
		if val == 2 {
			ret = append(ret, idx)
		}
	}

	return ret
}
