package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		res := strings.HasPrefix(url, "http://")
		if res == false {
			url = "http://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprint(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s", resp.Status)
	}
}
