package main

import (
    "fmt"
    "os"
    "strconv"
    "conv"
)


func main(){
    for _, arg := range os.Args[1:] {
        t, err := strconv.ParseFloat(arg, 64)
        if err != nil {
            fmt.Fprintf(os.Stderr, "cf: %v\n", err)
            os.Exit(1)
        }
        p := conv.Pound(t)
        k := conv.Kilogram(t)
        m := conv.Meter(t)
        f := conv.Foot(t)

        fmt.Printf("%s = %s, %s = %s, %s = %s, %s = %s\n", 
            p, conv.PToK(p), k, conv.KToP(k), m, conv.MToF(m), f, conv.FToM(f))

    }
}

