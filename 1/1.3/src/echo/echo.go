package echo

import (
	"strings"
)

func echo1(input []string) string {
	return strings.Join(input[1:], " ")
}

func echo2(input []string) string {
	s, sep := "", ""
	for _, arg := range input[1:] {
		s += sep + arg
		sep = " "
	}
	return s
}
