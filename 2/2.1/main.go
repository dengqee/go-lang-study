package main

import (
    "fmt"
    "tempconv"
)

func main(){
    fmt.Println(tempconv.CToF(tempconv.BoilingC))
    fmt.Println(tempconv.CToK(tempconv.AbsoluteZeroC))
    fmt.Println(tempconv.KToF(tempconv.AbsoluteZeroK))
}

