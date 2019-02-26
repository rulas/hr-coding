package main

import "fmt"

func fibbonacci(n int, mem []int64) int64 {
	if mem[n] != 0 {
		return mem[n]
	}

	if n == 0 {
		return 0
	}

	if n == 1 {
		mem[n] = 1
		return 1
	}

	result := fibbonacci(n-1, mem) + fibbonacci(n-2, mem)
	// fmt.Println("fib of " + string(n) + " is " + string(result))
	fmt.Printf("fib of %d is %d\n", n, result)
	mem[n] = result

	return result
}

func main() {
	const n int = 80

	mem := make([]int64, n+1)

	res := fibbonacci(n, mem)
	fmt.Println("result =", res)
}
