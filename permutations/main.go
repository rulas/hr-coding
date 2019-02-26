package main

import "fmt"

func swap(s []byte, l int, r int) {
	temp := s[r]
	s[r] = s[l]
	s[l] = temp
}

func permute(s []byte, l int, r int) {
	// fmt.Printf("Permuting %s from %d to %d\n", string(s), l, r)
	if l == r {
		fmt.Println(string(s))
	} else {
		for index := l; index <= r; index++ {
			swap(s, l, index)
			permute(s, l+1, r)
			swap(s, l, index)
		}
	}
}

func main() {
	s := "ABC"
	ss := make([]byte, len(s))
	n := len(s)
	for i := 0; i < n; i++ {
		ss[i] = s[i]
	}
	fmt.Println(string(ss))
	permute(ss, 0, n-1)
}
