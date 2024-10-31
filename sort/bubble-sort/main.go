// https://leetcode.com/problems/sort-colors/description/

package main

import "fmt"

func main() {
	arr := []int{2, 0, 2, 1, 1, 0}
	bubbleSort(arr)

	fmt.Println("Result:", arr)
}

func bubbleSort(arr []int) {
	end := len(arr) - 1

	for end > 0 {
		swapped := false
		for i := 0; i < end; i++ {
			if arr[i] > arr[i+1] {
				swapInArr(arr, i, i+1)
				swapped = true
			}
		}

		// if no swapped have occured, break the loop
		if !swapped {
			break
		}

		end--
	}
}

func swapInArr(arr []int, a, b int) {
	temp := arr[a]
	arr[a] = arr[b]
	arr[b] = temp
}
