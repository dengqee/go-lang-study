package main

import "fmt"

func main() {
	ch1, ch2 := make(chan string), make(chan string)
	go func() {
		for {
			ch1 <- "ping"
		}
	}()
	go func() {
		for {
			ch2 <- "pong"
		}
	}()
	for {
		select {
		case s := <-ch1:
			fmt.Println(s)
		case s := <-ch2:
			fmt.Println(s)
		}
	}

}
