package main

import "fmt"

func main(){
    f()
    fmt.Println("Hello, World!")
}
func f() {
    defer func() {
        if p:= recover();p!= nil {
            fmt.Printf("%v",p)
        }
    }()
    panic(1)
}

