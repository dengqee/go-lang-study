package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]
	counts := make(map[string]int)
	if len(files) == 0 {
		var s string
		countLines(s, os.Stdin, counts)
	} else {
		for _, arg := range files {
			counts := make(map[string]int)
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(arg, f, counts)
			f.Close()
		}
	}
}

func countLines(arg string, f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if counts[input.Text()] > 1 {
			fmt.Println(arg)
		}
	}
}
