package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	url := "https://www.sina.com/"
	words, images, err := CountWordsAndImages(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}
	fmt.Printf("%dwords,%dimages\n", words, images)
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	// switch n.Type {
	// case html.TextNode: fmt.Println("TextNode")
	// case html.DocumentNode: fmt.Println("DocumentNode")
	// case html.ElementNode: fmt.Println("ElementNode")
	// case html.CommentNode: fmt.Println("CommentNode")
	// case html.DoctypeNode: fmt.Println("DoctypeNode")
	// }
	if n.Type == html.ElementNode {
		if n.Data == "img" {
			images++
		}
	} else if n.Type == html.TextNode {
		scan := bufio.NewScanner(strings.NewReader(n.Data))
		scan.Split(bufio.ScanWords)
		for scan.Scan() {
			words++
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		// switch c.Type {
		// case html.TextNode: fmt.Println("TextNode")
		// case html.DocumentNode: fmt.Println("DocumentNode")
		// case html.ElementNode: fmt.Println("ElementNode")
		// case html.CommentNode: fmt.Println("CommentNode")
		// case html.DoctypeNode: fmt.Println("DoctypeNode")
		// }
		if n.Type == html.ElementNode {
			if n.Data == "img" {
				images++
			}
		} else if n.Type == html.TextNode {
			scan := bufio.NewScanner(strings.NewReader(n.Data))
			scan.Split(bufio.ScanWords)
			for scan.Scan() {
				words++
			}
		}

	}

	return

}
