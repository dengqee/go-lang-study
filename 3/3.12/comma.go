package main

import (
	"fmt"
	"os"
)

func main() {
	s1, s2 := os.Args[1], os.Args[2]
	fmt.Println(comma(s1, s2))
}

func comma(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	m1 := make(map[rune]uint32)
	m2 := make(map[rune]uint32)
	for _, s := range s1 {
		m1[s]++
	}
	for _, s := range s2 {
		m2[s]++
	}
	for s, c := range m1 {
		if m2[s] != c {
			return false
		}
	}
	return true

}
