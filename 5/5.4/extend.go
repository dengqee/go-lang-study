package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

var visitData = map[string]int{
	"a":      1,
	"img":    1,
	"link":   1,
	"script": 1,
}

func visit(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}
	if n.Type == html.ElementNode {
		if _, v := visitData[n.Data]; v {
			for _, a := range n.Attr {
				if a.Key == "href" {
					links = append(links, a.Val)
				}
			}
		}
	}
	links = visit(links, n.FirstChild)
	links = visit(links, n.NextSibling)
	return links
}
