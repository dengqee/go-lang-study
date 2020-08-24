package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		fmt.Printf("%x\n", sha256.Sum256([]byte(os.Args[1])))
	} else {
		f := os.Args[1]
		s := os.Args[2]
		c := []byte(s)
		switch f {
		case "384":
			fmt.Printf("%x\n", sha512.Sum384(c))
		case "512":
			fmt.Printf("%x\n", sha512.Sum512(c))
		default:
			fmt.Printf("%x\n", sha256.Sum256([]byte(os.Args[1])))

		}
	}

}
