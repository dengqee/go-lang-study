// 1.1 echo能够打印os.Args[0]
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])
}
