package main

import "fmt"

func findCells(arr []int, days int) []int {
	n := len(arr)
	dest := make([]int, n)
	copy(dest, arr)

	// iterate the days
	for i := 0; i < days; i++ {
		prevVal := 0
		// this is to iterate all the array, but the last (to simplify the loop)
		for j := 0; j < n-1; j++ {
			tmp := dest[j]
			dest[j] = prevVal ^ dest[j+1]
			prevVal = tmp
		}
		// this compute the last
		dest[n-1] = prevVal ^ 0
	}

	return dest
}

func main() {
	arr := []int{1, 0, 0, 0, 0, 1, 0, 0}
	newArr := findCells(arr, 1)
	fmt.Println(arr, "->", newArr)

	arr = []int{1, 1, 1, 0, 1, 1, 1, 1}
	newArr = findCells(arr, 2)
	fmt.Println(arr, "->", newArr)
}
