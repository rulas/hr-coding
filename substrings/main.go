package main

import "fmt"

func main() {
	s := "mom"
	n := len(s)
	alphs := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	fmt.Printf("[")
	for _, char := range alphs {
		fmt.Printf("%c ", char)
	}
	fmt.Printf("]\n")

	signatures := make(map[[]int]][]string)

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			ss := s[i : j+1]
			signature := uint32(0)
			signa := make([]byte)
			for _, char := range ss {
				index := uint32(char - 'a')
				signature |= 0x80000000 >> index
			}
			fmt.Printf("signature=%032b, string=%s\n", signature, ss)
			signatures[signature] = append(signatures[signature], ss)
			// fmt.Println(ss)
			// fmt.Println(signature)
			// fmt.Printf("%d, %d\n", i, j)
		}
	}
	fmt.Println(signatures)
}
