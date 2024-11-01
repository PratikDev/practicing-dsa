package main

import "fmt"

func main() {
	arr := []int{34, -20, 17, 5, 19, 23, 51, -33}
	insertionSort(arr)
	fmt.Println("Result:", arr)
}

func insertionSort(arr []int) {
	for start := 1; start <= len(arr)-1; start++ {
		for i := start; i > 0; i-- {
			if arr[i] < arr[i-1] {
				swap(arr, i, i-1)
			} else {
				break
			}
		}
	}
}

func swap(arr []int, a, b int) {
	temp := arr[a]
	arr[a] = arr[b]
	arr[b] = temp
}
