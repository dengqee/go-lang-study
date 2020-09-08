// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 241.

// Crawl2 crawls web links starting with the command-line arguments.
//
// This version uses a buffered channel as a counting semaphore
// to limit the number of concurrent calls to links.Extract.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"../../../../gopl.io/ch5/links"
)

//!+sema
// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)
var depth = flag.Int("depth", -1, "limite find depth")

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token

	if err != nil {
		log.Print(err)
	}
	return list
}

//!-sema

//!+
type linkdepth struct {
	links []string
	depth int
}

func main() {
	flag.Parse()
	worklist := make(chan linkdepth)
	var n int // number of pending sends to worklist
	fmt.Println(*depth)

	// Start with the command-line arguments.
	n++
	go func() { worklist <- linkdepth{links: os.Args[1:], depth: 0} }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		if *depth != -1 && list.depth >= *depth {
			fmt.Println(*depth)
			break
		}
		for _, link := range list.links {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- linkdepth{crawl(link), list.depth + 1}
				}(link)
			}
		}
	}
}

//!-
