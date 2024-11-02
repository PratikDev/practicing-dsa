// https://leetcode.com/problems/find-the-duplicate-number/description/

package main

import "fmt"

func main() {
	arr := []int{5, 4, 2, 2, 1, 3, 6}
	result := findDuplicate(arr)
	fmt.Println(result)
}

func findDuplicate(arr []int) int {
	n := len(arr)
	ret := make([]int, n+1)

	for _, v := range arr {
		if ret[v] == 0 {
			ret[v] = 1
		} else {
			return v
		}
	}

	return -1
}
