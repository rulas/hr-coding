package main

import (
	"fmt"
)

type Pos struct {
	x int
	y int
}

func getShortes() {

}

func findShortestDistance(lot [][]int, start []int) int {
	fmt.Println(lot, start)
	x, y := start[0], start[1]

	w[x][y] = 0
	

	visited := make(map[Pos]bool)
	visited[Pos{0, 3}] = true
	fmt.Println(visited)
	return 0
}

func main() {
	input := [][]int{
		{1, 1, 1, 1},
		{0, 1, 1, 1},
		{0, 1, 0, 1},
		{1, 1, 9, 1},
		{0, 0, 1, 1},
	}

	retVal := findShortestDistance(input, []int{0, 0})
	fmt.Println("shortest path: ", retVal)
}
