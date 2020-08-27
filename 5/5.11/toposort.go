package main

import (
	"fmt"
	"sort"
    "os"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
    "linear algebra": {"calculus"},
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
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
            for _,learned := range order {
                if learned == item {
                    fmt.Errorf("%s成环",item)
                    os.Exit(1)
                }
            }
			if seen[item]==false {
				order = append(order, item)
				visitAll(m[item])
				seen[item] = true
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}
