/*
    Write a simple program to compute the tax on an order
    amount. The program should prompt for the order amount
    and the state. If the state is "WI" then the order must
    be charged 5.5% tax. The program should display the subtotal,
    tax and total for Wisconsin residents but display just the total
    for non residents

    Example:

    What is the order amount? 10
    What is the State? WI
    The subtotal is $10.00
    The tax is $0.55
    The total is $10.55

    OR

    What is the order amount? 10
    What is the State? MN
    The total is $10.00

 */

package main

import (
    "fmt"
    "strings"
)

func GenerateOutput(state string, amount, tax, total float64) string {
    var line1,line2 string
    if tax != 0 {
        line1 = fmt.Sprintf("The subtotal is %.2f\n", amount)
        line2 = fmt.Sprintf("The tax in %s is %.2f\n",state, tax)
    }
    line3 := fmt.Sprintf("The total is %.2f\n", total)
    return line1 + line2 + line3
}

func CalculateTax(amount, rate float64) float64 {
    return amount * rate / 100
}

func totaller(nums ...float64) float64 {
    var total float64
    for _, v := range nums {
        total += v
    }
    return total
}

func main(){
    // Declare variables
    var (
        amount float64
        state string
        outputStr string
        taxRate float64
        total float64
    )
    // Get user input
    fmt.Print("What is the order amount? ")
    fmt.Scanf("%f ", &amount)

    fmt.Print("What is the state? ")
    fmt.Scanf("%s ", &state)

    // Convert and clean user input
    state = strings.TrimSpace(state)

    // Determine tax status
    if state == "WI" {
        taxRate = 0.055
    }

    tax := CalculateTax(amount, taxRate)
    total = totaller(amount, tax)
    outputStr = GenerateOutput(state, amount, tax, total)

    // Print output
    fmt.Println(outputStr)
}