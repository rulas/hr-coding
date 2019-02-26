package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

//  0 1 2 3 4 5 6 7
// mid[0] = 3
//
func binarySearch(arr []int, num int) int {
	low := 0
	high := len(arr)
	isFound := false
	pos := -1

	for !isFound {
		middle := (high-low)/2 + low
		log.Println("low=", low, "high=", high, "middle=", middle)

		if low > high {
			return -1
		}

		if num > arr[middle] {
			low = middle + 1
		} else if num < arr[middle] {
			high = middle - 1
		} else if arr[middle] == num {
			return middle
		}
	}

	return pos
}

func main() {
	arr := []int{1, 6, 22, 25, 94, 123, 138, 200, 223, 250, 500}
	fmt.Println("arr=", arr)
	num, ok := strconv.Atoi(os.Args[1])
	if ok != nil {
		log.Fatal("kaboooom")
	}

	pos := binarySearch(arr, num)
	if pos < 0 {
		fmt.Println(num, "not found")
	} else {
		fmt.Println(num, "found at", pos)
	}
}
