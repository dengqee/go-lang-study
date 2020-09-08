package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func main() {
	//addr := os.Args[1]
	addr := "localhost:8088"
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Errorf("%v", err)
		os.Exit(1)
	}
	defer conn.Close()
	info := bufio.NewScanner(conn)
	info.Scan()
	pwd := info.Text()
	fmt.Printf("ftp:%s>", pwd)
	cmdline := bufio.NewScanner(os.Stdin)
	var cmd string
	for {
		cmdline.Scan()
		cmd = cmdline.Text()
		_, err := io.WriteString(conn, cmd+"\n")
		if err != nil {
			fmt.Errorf("%v\n", err)
			fmt.Printf("ftp:%s>", pwd)
			continue
		}
		cmdInfo := strings.Fields(cmd)
		switch cmdInfo[0] {
		case "cd":
			info.Scan()
			if info.Text() == "err" {
				fmt.Println("erro!")
				fmt.Printf("ftp:%s>", pwd)
				continue
			}
			pwd = info.Text()
		case "ls":
			info.Scan()
			if info.Text() == "err" {
				fmt.Println("erro!")
				fmt.Printf("ftp:%s>", pwd)
				continue
			}
			for info.Text() != "EOF" {
				fmt.Println(info.Text())
				info.Scan()
			}
		case "exit":
			fmt.Println("exit!")
			os.Exit(0)
		}
		fmt.Printf("ftp:%s>", pwd)
	}

}
