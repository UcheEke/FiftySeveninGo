/*
    Password Validation

    Create a simple program that validates user login credentials
    The program must prompt the user for a username and
    password. The program should compare the password given by the user
    to a known password. If the password matches the program should display
    "Welcome" If it doesn't match the program should display
    "I don't know you"
 */

package main

import (
    "fmt"
)

const (
    internalPassword = "opensesame"
)

func checkPassword(password string) bool {
    return password == internalPassword
}

func main() {
    var password string

    fmt.Print("What is the password? ")
    fmt.Scanf("%s\n", &password)
    if checkPassword(password) {
        fmt.Println("Welcome!")
    } else {
        fmt.Println("I don't know you!")
    }
}