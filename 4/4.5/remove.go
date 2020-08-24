package main

import "fmt"

func main() {
	s := []string{"ab", "cd", "cd", "cd", "ab"}
	fmt.Println(remove(s))
}

func remove(s []string) []string {
	l := len(s)
	if l < 2 {
		return s
	}
	var a, b string
	for i := 0; i < l-1; i++ {
		a, b = s[i], s[i+1]
		if a == b {
			copy(s[i:], s[i+1:])
			l--
			i--
		}
	}
	return s[0:l]

}
