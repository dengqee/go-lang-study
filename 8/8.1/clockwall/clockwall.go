package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

type timetTable struct {
	addr string
	t    string
}

func (tt *timetTable) Write(p []byte) (n int, err error) {
	tt.t = string(p)
	fmt.Printf("%s-%s", tt.addr, tt.t)
	n = len(tt.t)
	return n, nil
}

func main() {
	addrs := os.Args[1:]
	//addrs := []string{"localhost:8010", "localhost:8020"}
	for _, addr := range addrs {
		conn, _ := net.Dial("tcp", addr)
		defer conn.Close()
		go clockCopy(conn, addr)
	}
	for {
	}
}
func clockCopy(src io.Reader, addr string) {
	tt := timetTable{addr: addr}
	fmt.Println(addr)
	io.Copy(&tt, src)
}
