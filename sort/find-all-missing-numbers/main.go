package main

import "fmt"

func main() {
	arr := []int{1, 1}
	result := findAllMissingNumbers(arr)
	fmt.Println(result)
}

func findAllMissingNumbers(arr []int) []int {
	// sort it using cyclic sort
	for i := 0; i < len(arr); {
		value := arr[i]

		if value == arr[value-1] {
			i++
		} else {
			// swap
			arr[i] = arr[value-1]
			arr[value-1] = value
		}
	}

	var results []int
	for i, value := range arr {
		if value != i+1 {
			results = append(results, i+1)
		}
	}

	return results
}
