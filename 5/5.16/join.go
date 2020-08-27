package main

import "fmt"

func main(){
    
    fmt.Println(join("a"," ","b"))
}

func join(ss ...string)string {
    var res string
    for _,s := range ss {
        res += s
    }
    return res
}

