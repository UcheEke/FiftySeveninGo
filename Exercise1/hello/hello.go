package main

import (
    "fmt"
)

func main(){
    var name string
    var result string
    fmt.Print("What is your name? ")
    fmt.Scanf("%s", &name)

    result = fmt.Sprintf("Hello, %s! Nice to meet you!", name)
    fmt.Println(result)
}