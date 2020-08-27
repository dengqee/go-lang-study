package main

import "fmt"

func main() {
	s := "abcdfoef"
	fmt.Println(expand(s, f))
}

func expand(s string, f func(string) string) string {
	i := 0
	for i = 0; i < len(s)-3; i++ {
		if s[i:i+3] == "foo" {
			break
		}
	}
	if i < len(s)-3 {
		return s[:i] + f("foo") + s[i+3:]
	}
	return s
}

func f(s string) string {
	if s == "foo" {
		return "fee"
	}
	return "fcc"
}
