package main

import "fmt"

func main() {
	arr := []int{5, 2, 3, 1}
	selectionSort(arr)

	fmt.Println("Result:", arr)
}

func selectionSort(arr []int) {
	for end := len(arr) - 1; end > 0; end-- {
		highestIndex := findHighestIndex(arr[:end+1])
		swapInArray(arr, highestIndex, end)
	}
}

func findHighestIndex(arr []int) int {
	highestIndex := 0
	for idx, value := range arr {
		if value > arr[highestIndex] {
			highestIndex = idx
		}
	}
	return highestIndex
}

func swapInArray(arr []int, a, b int) {
	temp := arr[a]
	arr[a] = arr[b]
	arr[b] = temp
}
