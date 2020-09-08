// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 224.

// Reverb2 is a TCP server that simulates an echo.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}
func closEcho(c net.Conn) {
	fmt.Println("close")
	c.Close()
}

var t int = 10

func timer() {
	for t > 0 {
		time.Sleep(1 * time.Second)
		t--
	}
	ch <- struct{}{}
}

var ch chan struct{} = make(chan struct{})

//!+
func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	go func() {
		for input.Scan() {
			t = 10
			go echo(c, input.Text(), 1*time.Second)
		}
	}()
	go timer()
	select {
	case <-ch:
		c.Close()
		return
	}
	// NOTE: ignoring potential errors from input.Err()
	closEcho(c)
}

//!-

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
