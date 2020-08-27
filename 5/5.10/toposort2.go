package main

import (
	"fmt"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structtures", "computer organization"},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}
func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(map[string][]string)
	visitAll = func(items map[string][]string) {
		for item, pre := range items {
			if seen[item] == false {
				seen[item] = true
				preMap := make(map[string][]string)
				for _, i := range pre {
					preMap[i] = m[i]
				}
				visitAll(preMap)
				order = append(order, item)
			}
		}
	}
	visitAll(m)
	return order
}
