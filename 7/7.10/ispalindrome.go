package main

import (
	"fmt"
	"sort"
)

type str []byte

func (p str) Len() int           { return len(p) }
func (p str) Less(i, j int) bool { return p[i] < p[j] }
func (p str) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func isPalindrome(p sort.Interface) bool {
	for i, j := 0, p.Len()-1; i < j; {
		if p.Less(i, j) || p.Less(j, i) {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	p := str("abcba")
	fmt.Println(isPalindrome(p))
}
