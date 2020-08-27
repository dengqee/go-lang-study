package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"os"

	"golang.org/x/net/html"
)

type NodeCount map[string]int

const templ = `
--------------------------------------
|       HTML NODE
{{- range $key, $value := .}}
|   Type: {{$key | printf "%-15s" }} Count: {{$value | printf "%-4d" }} |
{{- end }}
`

func (nc NodeCount) Fprint(w io.Writer) {
	t := template.Must(template.New("excape").Parse(templ))
	if err := t.Execute(w, nc); err != nil {
		log.Fatal(err)
	}
}
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	nodeCount := NodeCount{}
	count(&nodeCount, doc)
	nodeCount.Fprint(os.Stdout)
}

func count(nc *NodeCount, n *html.Node) {
	if n == nil {
		return
	}
	if n.Type == html.ElementNode {
		(*nc)[n.Data]++
	}
	count(nc, n.FirstChild)
	count(nc, n.NextSibling)
}
