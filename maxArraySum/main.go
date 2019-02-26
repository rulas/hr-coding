package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the maximumSum function below.
func maximumSum(a []int64, m int64) int64 {
	size := len(a)
	// fmt.Println(a)

	// curr := int64(0)
	// prefix := make([]int64, len(a))

	// for i := 0; i < size; i++ {
	// 	curr = (a[i]%m + curr) % m
	// 	prefix[i] = curr
	// }

	max := int64(0)

	for l := 0; l < size; l++ {
		max_so_far := int64(0)
		for r := l + 1; r <= size; r++ {
			curSum := int64(0)
			for i := l; i < r; i++ {
				curSum += a[i]
				// fmt.Printf("summing %d\n", a[i])
			}
			mod := curSum % m
			// fmt.Printf("summming from %d to %d = %d\n", l, r, mod)

			if mod > max {
				max = mod
			}
		}
	}

	return max
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		nm := strings.Split(readLine(reader), " ")

		nTemp, err := strconv.ParseInt(nm[0], 10, 64)
		checkError(err)
		n := int32(nTemp)

		m, err := strconv.ParseInt(nm[1], 10, 64)
		checkError(err)

		aTemp := strings.Split(readLine(reader), " ")

		var a []int64

		for i := 0; i < int(n); i++ {
			aItem, err := strconv.ParseInt(aTemp[i], 10, 64)
			checkError(err)
			a = append(a, aItem)
		}

		result := maximumSum(a, m)

		fmt.Fprintf(writer, "%d\n", result)
	}

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
