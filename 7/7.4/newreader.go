package main

import (
	"fmt"
	"golang.org/x/net/html"
	"strings"
)

func NewReader(s string) (*html.Node, error) {
	read := strings.NewReader(s)
	doc, err := html.Parse(read)
	return doc, err
}
func main() {
	doc, _ := NewReader("https://www.sina.com/")

	fmt.Println("http:")
	fmt.Println(doc)
}
