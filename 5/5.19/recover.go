package main

import "fmt"

func main() {
	fmt.Println(f())
}
func f() (r int) {
	defer func() {
		if p := recover(); p != nil {
			r = 1
		} else {
			r = 2
		}

	}()
	panic(1)
}
