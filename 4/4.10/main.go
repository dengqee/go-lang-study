package main

import (
    "fmt"
    "log"
    "os"

    "github"
)

func main(){
    result, err := github.SearchIssues(os.Args[1:])
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%d issues:\n", result.TotalCount)
    now := time.Now().Unix()
    preMonth := now - 30*3600
    preYear := now - 365*24*3600
    var notMonth []*github.Issue
    var notYear []*github.Issue
    var overYear []*github.Issue


    for _, item := range result.Items {
        createTime := item.CreateAt.Unix()
        if createTime > preMonth {
            notMonth = append(notMonth,item)
            continue
        }
        if createTime < preMonth && createTime > preYear {
            notYear = append(notYear,item)
            continue
        }
        overYear = append(overYear, item)
    }

    fmt.Println("notMonth:")
    for _,item := range notMonth {
        fmt.Printf("#%-5d %9.9s %.55s time:%s\n",item.Number, item.User.Login,item.CreateAt)
    }
    fmt.Println("notYear:")
    for _,item := range notYear {
        fmt.Printf("#%-5d %9.9s %.55s time:%s\n",item.Number, item.User.Login,item.CreateAt)
    }
    fmt.Println("overYear:")
    for _,item := range overYear {
        fmt.Printf("#%-5d %9.9s %.55s time:%s\n",item.Number, item.User.Login,item.CreateAt)
    }
}

