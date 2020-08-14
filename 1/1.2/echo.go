//1.2 打印每个参数的索引和值
package main

import (
	"fmt"
	"os"
)

func main() {
	for id, arg := range os.Args[1:] {
		fmt.Println(id+1, " ", arg)
	}
}
