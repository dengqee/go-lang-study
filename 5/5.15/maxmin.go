package main

import "fmt"
const (
    INT_MAX = int(^uint(0)>>1)
    INT_MIN = ^INT_MAX
)
func main(){
    fmt.Printf("%d\n",max())
    fmt.Printf("%d\n",max(0,1,2,3,4))
    fmt.Printf("%d\n",min(0,1,2,3,4))
}

func max(vals ...int) int {
    res := INT_MIN
    if len(vals)<1 {
        fmt.Println("no input");
        return res
    }
    for _, val := range vals {
        if res< val {
            res = val
        }
    }
    return res
}

func min(vals ...int) int {
    res := INT_MAX
    if len(vals)<1 {
        fmt.Errorf("no input");
        return res
    }
    for _, val := range vals {
        if res> val {
            res = val
        }
    }
    return res
}

