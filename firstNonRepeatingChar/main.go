package main

import (
	"fmt"
)

func getCharCount(s string) []byte {
	charCount := make([]byte, 26)
	for _, v := range s {
		charCount[v-'a']++
	}
	fmt.Println(&charCount[0])
	return charCount
}

func firstNotRepeatingCharacter(s string) string {

	charCount := getCharCount(s)
	fmt.Println(&charCount[0])

	for _, v := range s {
		if charCount[v-'a'] == 1 {
			return string(v)
		}
	}

	return "_"
}

func main() {

	fmt.Println("aaa" > "bbb")

	a := "abacabad"
	s := firstNotRepeatingCharacter(a)
	fmt.Println(s)

	m := [][]int{{2, 3, 4},
		{5, 6, 7}}
	fmt.Println(len(m), cap(m), m)

}
