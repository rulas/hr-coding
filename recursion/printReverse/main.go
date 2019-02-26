package main

import "fmt"

func PrintReverse(s string) {
	printReverse(s, 0)
}

func printReverse(s string, pos int) {
	if len(s) == 0 || pos >= len(s) {
		return
	}

	printReverse(s, pos+1)
	fmt.Printf("%c", s[pos])

}

func main() {
	s := "Hello world"
	PrintReverse(s)
}
