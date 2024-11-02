package main

import "fmt"

func main() {
	arr := []int{1, 2}
	cyclicSort(arr)
	fmt.Println("Result:", arr)
}

func cyclicSort(arr []int) {
	i := 0
	for i < len(arr) {
		value := arr[i]

		if i+1 == value {
			i++
		} else {
			// swap
			arr[i] = arr[value-1]
			arr[value-1] = value
		}
	}
}

func cyclicSortAlt(arr []int) {
	last := len(arr) - 1
	for i := 0; i <= last; i++ {
		arr[i] = i + 1
	}
}
