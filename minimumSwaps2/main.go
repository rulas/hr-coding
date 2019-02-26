package main

import (
	"fmt"
)

// Complete the minimumSwaps function below.
func minimumSwaps(arr []int32) int32 {
	var swapCount int32

	for index := 0; index < len(arr)-1; {

		if val := arr[index]; val == int32(index+1) {
			index++
		} else {
			arr[index], arr[val-1] = arr[val-1], arr[index]
			swapCount++
		}
	}

	return swapCount
}

func main() {

	arr := []int32{1, 3, 5, 2, 4, 6, 7}
	res := minimumSwaps(arr)

	fmt.Println(res)
}
