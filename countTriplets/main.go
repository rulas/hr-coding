package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// Complete the countTriplets function below.
func countTriplets(arr []int64, r int64) int64 {
	log.Println("input: ", arr)
	n := len(arr)
	left := make(map[int64]int)
	right := make(map[int64]int)

	// fill the maps with initial values
	left[arr[0]] = 1
	for i := 2; i < n; i++ {
		right[arr[i]]++
	}
	log.Println("original maps:", left, right)

	count := 0
	for i := 1; i < n-1; i++ {
		curr := arr[i]
		lcount := left[curr/r]
		rcount := right[curr*r]
		log.Println("pre-values:", i-1, i, i+1, left, right)
		count += lcount * rcount
		left[curr]++
		right[curr]--
		log.Println("post-values:", i-1, i, i+1, left, right)

	}

	log.Println(count)
	return i 
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nr := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(nr[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	r, err := strconv.ParseInt(nr[1], 10, 64)
	checkError(err)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int64

	for i := 0; i < int(n); i++ {
		arrItem, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arr = append(arr, arrItem)
	}

	ans := countTriplets(arr, r)

	fmt.Fprintf(writer, "%d\n", ans)

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
