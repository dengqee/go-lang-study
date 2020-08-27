package main

import (
    "fmt"
    "os"
    "./links"
    "log"
)
var f *os.File
func main(){
    f,_ = os.Create("./url")
    breadthFirst(crawl,os.Args[1:])
    f.Close()
}

func breadthFirst(f func(item string) []string, worklist []string) {
    seen := make(map[string]bool)
    for len(worklist) > 0 {
        items := worklist
        worklist = nil 
        for _, item := range items {
            if seen[item] == false {
                seen[item] = true
                worklist = append(worklist, f(item)...)
            }
        }
    }
}

func crawl(url string) []string {
    fmt.Println(url)
    list, err := links.Extract(url)
    if err != nil {
        log.Print(err)
    }
    for _, urli := range list {
        f.WriteString(urli+"\n")
    }
    return list
}

