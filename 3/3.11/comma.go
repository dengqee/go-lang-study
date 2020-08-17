package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	s := os.Args[1]
	s = comma(s)
	fmt.Println(s)
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	var buf bytes.Buffer
	if s[0] == '+' || s[0] == '-' {
		buf.WriteString(string(s[0]))
		s = s[1:]
	}
	pos := 0
	for ; pos < n; pos++ {
		if s[pos] == '.' {
			break
		}
	}
	s1 := s[:pos]
	n1 := len(s1)

	for i := 0; i < n1; i++ {
		buf.WriteString(string(s1[i]))
		if (i < n1-1) && ((i+1)%3 == n1%3) {
			buf.WriteString(",")
		}
	}
	buf.WriteString(s[pos:])
	return buf.String()
}
