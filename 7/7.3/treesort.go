package main

import (
    "fmt"
    "bytes"
)

type Stringer interface {
	String() string
}

type tree struct {
	value       int
	left, right *tree
}

func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
	fmt.Println(root.String())
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}

	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func (t *tree) String() string {
	var s []int
	s=f(t, s)
    var buf bytes.Buffer
    buf.WriteByte('{')
    for _,v := range s {
        fmt.Fprintf(&buf, " %d", v)
    }
    buf.WriteByte('}')

	return buf.String()
}
func f(t *tree, s []int) []int{
	if t != nil {
		s=f(t.left, s)
		s = append(s, t.value)
		s=f(t.right, s)
	}
    return s

}

func main() {
	a := []int{4, 5, 3, 5, 7, 1}
	Sort(a)
}
