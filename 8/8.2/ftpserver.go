package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
)

type ftpServer struct {
	pwd  string
	conn net.Conn
}

func (ftp *ftpServer) handleConn(c net.Conn) {
	defer c.Close()
	ftp.conn = c
	io.WriteString(c, ftp.pwd+"\n")
	input := bufio.NewScanner(c)
	var cmd []string
	for {
		//fmt.Println("25:for")
		input.Scan()
		//fmt.Println("27:Scan")
		cmdStr := input.Text()
		if cmdStr == "exit" {
			return
		}
		cmd = strings.Fields(cmdStr)
		switch cmd[0] {
		case "cd":
			ftp.handleCd(cmd)

		case "ls":
			ftp.handleLs(cmd)
		}
	}
}
func (ftp *ftpServer) handleCd(cmd []string) {
	//var cmdLine string
	//for _, c := range cmd {
	//	cmdLine += c + " "
	//}
	//fmt.Println(cmdLine)
	if len(cmd) < 2 {
		io.WriteString(ftp.conn, "err\n")
		return
	}
	err := os.Chdir(cmd[1])
	if err != nil {
		io.WriteString(ftp.conn, "err\n")
		return
	}

	pwd, _ := os.Getwd()
	ftp.pwd = pwd
	fmt.Println(pwd)
	io.WriteString(ftp.conn, ftp.pwd+"\n")
}
func (ftp *ftpServer) handleLs(cmd []string) {
	if len(cmd) >= 2 { //only "ls"
		io.WriteString(ftp.conn, "err\n")

	}
	out, err := exec.Command("ls").Output()
	if err != nil {
		io.WriteString(ftp.conn, "err\n")
	}
	io.WriteString(ftp.conn, string(out))
	fmt.Fprintf(ftp.conn, "%v\n", io.EOF)

}

func main() {
	addr := "localhost:8088"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Println("conneted")
		pwd, _ := os.Getwd()
		fmt.Println(pwd)
		ftp := ftpServer{pwd: pwd}
		go ftp.handleConn(conn)
	}

}
