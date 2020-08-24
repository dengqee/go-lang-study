package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := "a  ac  d s"

	fmt.Println(remove(s))
}
func remove(a string) string {
	s := []byte(a)

	l := len(s)
	if l <= 1 {
		return a
	}
	for i := 0; i < l-1; i++ {
		if unicode.IsSpace(rune(s[i])) && unicode.IsSpace(rune(s[i+1])) {
			copy(s[i:], s[i+1:])
			l--
			i--
		}
	}
	res := string(s[:l])
	return res
}
