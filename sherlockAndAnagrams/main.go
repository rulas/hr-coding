package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func getSignature(s string) string {
	charCounter := make([]byte, 26)
	for i := 0; i < len(s); i++ {
		index := s[i] - 'a'
		charCounter[index]++
	}

	var str strings.Builder
	for _, char := range charCounter {
		str.WriteString(fmt.Sprintf("%d", char))
		str.WriteString(";")
	}

	return str.String()
}

// Complete the sherlockAndAnagrams function below.
func sherlockAndAnagrams(s string) int32 {
	n := len(s)
	signatures := make(map[string][]string)

	// get all the substring
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			ss := s[i : j+1]
			signature := getSignature(ss)
			signatures[signature] = append(signatures[signature], ss)
		}
	}

	// find all the signatures
	for k, v := range signatures {
		if len(v) == 1 {
			delete(signatures, k)
		}
	}

	sum := 0
	for k, v := range signatures {
		n = len(v)
		sum += n * (n - 1) / 2
		fmt.Println(k, v)
	}
	// fmt.Println(signatures)

	return int32(sum)
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
		s := readLine(reader)

		result := sherlockAndAnagrams(s)

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
