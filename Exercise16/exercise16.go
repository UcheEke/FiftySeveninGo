/*
   Simple tests for the legal driving limit
 */

package main

import (
    "fmt"
)

func canDrive(age int) bool {
    return age > 16
}

func main(){
    var age int
    fmt.Print("What is your age? ")
    fmt.Scanf("%d ",&age)

    if canDrive(age) {
        fmt.Println("You are legal to drive")
    } else {
        fmt.Println("You are not at the legal age limit yet")
    }
}