package main

import (
	"bytes"
	"fmt"
)

func main() {
	s := "1234567"
	s = comma(s)
	fmt.Println(s)
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		buf.WriteString(string(s[i]))
		if (i < n-1) && ((i+1)%3 == n%3) {
			buf.WriteString(",")
		}
	}
	return buf.String()
}
