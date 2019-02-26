package main

import "fmt"

// Complete the minimumBribes function below.
func minimumBribes(q []int32) {
	// fmt.Println(q)
	numBribes := int32(0)
	hasSwap := true

	for hasSwap {
		hasSwap = false
		for i := 0; i < len(q)-1; i++ {
			if int(q[i])-2 > i+1 {
				fmt.Println("Too chaotic")
				return
			}
			if q[i] > q[i+1] {
				q[i], q[i+1] = q[i+1], q[i]
				numBribes++
				hasSwap = true
			}
		}
	}

	fmt.Println(numBribes)
}

func main() {

	q := []int32{1, 2, 5, 3, 7, 8, 6, 4}

	minimumBribes(q)
}
