// https://leetcode.com/problems/missing-number/description/

package main

import "fmt"

func main() {
	arr := []int{1, 2}
	fmt.Println(find(arr))
}

func find(arr []int) int {
	for i := 0; i < len(arr); {
		value := arr[i]

		if i == value || value == len(arr) {
			i++
		} else {
			arr[i] = arr[value]
			arr[value] = value
		}
	}

	for i := 0; i < len(arr); i++ {
		if i != arr[i] {
			return i
		}
	}

	return len(arr)
}
