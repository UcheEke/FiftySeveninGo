/*
    Create a program that generates multiplication tables for numbers 0 through 12
 */

package main

import (
    "fmt"
)

func main(){
    fmt.Println("Multiplication tables")
    for i := 0; i < 13; i++ {
        fmt.Printf("\nThe %d times table\n", i)
        for j := 0; j < 13; j++ {
            fmt.Println(i,"x",j,"=",i*j)
        }
    }
}