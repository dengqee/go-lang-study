package main

import (
	"bufio"
	"fmt"
	"strings"
)

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanWords)
	// scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		*c++
	}
	return len(p), nil

}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		*c++
	}
	return len(p), nil

}
func main() {
	var c WordCounter
	c.Write([]byte("hello world"))
	fmt.Println(c)

	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)

	var l LineCounter
	l.Write([]byte("hello\n world\n"))
	fmt.Println(l)

	l = 0
	fmt.Fprintf(&l, "hello\n world\n %s\n", name)
	fmt.Println(l)
}
