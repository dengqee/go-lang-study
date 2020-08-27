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
	id := "href"
	node := ElementByID(doc, id)
	if node != nil {
		fmt.Printf("%s<%s", "", node.Data)
		for _, a := range node.Attr {
			fmt.Printf(" %s=\"%s\"", a.Key, a.Val)
		}
		fmt.Printf(">\n")
	} else {
		fmt.Println("nil")
	}
}

func forEachNode(n *html.Node, pre func(n *html.Node, id string) bool, id string) *html.Node {
	if pre != nil {
		if pre(n, id) == true {
			return n
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if node := forEachNode(c, pre, id); node != nil {
			return node
		}
	}
	return nil
}
func ElementByID(doc *html.Node, id string) *html.Node {
	return forEachNode(doc, startElement, id)

}

func startElement(n *html.Node, id string) bool {
	if n == nil {
		return false
	}
	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			if a.Key == id {
				return true
			}
		}
		return false
	}
	return false
}
