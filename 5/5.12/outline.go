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
    forEachNode(doc, nil, nil)
}

func forEachNode(n *html.Node, pre, post func(n *html.Node,depth *int)) {
    pre = func(n *html.Node, depth *int) {
        if n == nil {
            return
        }
        if n.Type == html.ElementNode {
            fmt.Printf("%*s<%s", *depth*2, "", n.Data)
            for _,a := range n.Attr {
                fmt.Printf(" %s=\"%s\"", a.Key, a.Val)
            }
            fmt.Printf(">\n")
            (*depth)++
        }
    }
    post = func(n *html.Node, depth *int){
        if n.Type == html.ElementNode {
            (*depth)--
            fmt.Printf("%*s</%s>\n", *depth*2, "", n.Data)
        }
        
    }
    pre(n,&depth)
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        forEachNode(c, pre, post)
    }

    post(n,&depth)

}

var depth int
