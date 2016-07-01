/*
    Write a program that converts currency. Specifically,
    convert euros to US dollars. Prompt for the amount
    of money in euros you have. Print out the new amount in US
    dollars.

 */

package main

import (
    "fmt"
    "math"
)

func euros2dollars(euros, d2p, p2e float64) float64 {
    return euros * p2e / d2p
}

func fix(val, sigf float64) float64 {
    t := math.Pow(10, sigf)
    return math.Floor(val * t) / t
}

func main() {
    var (
        d2p float64
        p2e float64
        euros float64
    )
    // Prompt for user input
    fmt.Println("Euro to US $ Conversion:")
    fmt.Print("Amount in Euro: ")
    fmt.Scanf("%f ", &euros)

    fmt.Print("Exchange rate $/£: ")
    fmt.Scanf("%f ", &d2p)

    fmt.Print("Exchange rate £/e: ")
    fmt.Scanf("%f ", &p2e)

    // Calculate dollar amount
    dollars := euros2dollars(euros, d2p, p2e)

    fmt.Printf("Amount in US Dollars: %.2f\n", dollars)
}