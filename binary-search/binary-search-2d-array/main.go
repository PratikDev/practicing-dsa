package main

import "fmt"

func main() {
	matrix := [][]int{
		{1},
	}
	target := 3
	result := binarySearch2DArray(matrix, target)
	fmt.Println("Result:", result)
}

func binarySearch2DArray(arr [][]int, target int) bool {
	row := 0
	col := len(arr[0]) - 1

	for row < len(arr) && col >= 0 {
		if arr[row][col] == target {
			return true
		}

		if arr[row][col] < target {
			row++
		} else {
			col--
		}
	}

	return false
}
