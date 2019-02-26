package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the arrayManipulation function below.
func arrayManipulation(n int32, queries [][]int32) int64 {
	arr := make([]int64, int(n)+1)
	for _, v := range queries {
		for i := v[0]; i <= v[1]; i++ {
			arr[i] = arr[i] + int64(v[2])
		}
	}

	max := int64(0)
	for _, v := range arr {
		if v > max {
			max = v
		}
	}

	return max
}

func arrayManipulation2(n int32, queries [][]int32) int64 {
	dArr := make([]int64, int(n)+2)
	max := int64(0)

	for _, query := range queries {
		l, r, x := query[0], query[1], query[2]
		dArr[l] += int64(x)
		dArr[r+1] -= int64(x)
		fmt.Println(dArr)
	}

	x := int64(0)
	size := len(dArr)
	for i := 0; i < size-1; i++ {
		x += dArr[i]
		if x > max {
			max = x
		}
	}

	return max
}

func main() {

	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create("output.txt")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nm := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nm[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(nm[1], 10, 64)
	checkError(err)
	m := int32(mTemp)

	var queries [][]int32
	for i := 0; i < int(m); i++ {
		queriesRowTemp := strings.Split(readLine(reader), " ")

		var queriesRow []int32
		for _, queriesRowItem := range queriesRowTemp {
			queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
			checkError(err)
			queriesItem := int32(queriesItemTemp)
			queriesRow = append(queriesRow, queriesItem)
		}

		if len(queriesRow) != int(3) {
			panic("Bad input")
		}

		queries = append(queries, queriesRow)
	}

	result := arrayManipulation2(n, queries)

	fmt.Fprintf(writer, "%d\n", result)
	fmt.Println(result)

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
