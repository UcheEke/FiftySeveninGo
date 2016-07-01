package main

import (
    "../../../Projects/FiftySeven/ReflectorPrompt"
    "math"
    "fmt"
)

func taxRate(state string) float64 {
    var tr float64
    switch state {
    case "WI":
        tr = countyRate() + 0.05
    case "IL":
        tr = 0.08
    }

    return tr
}

func countyRate() float64 {
    var county string
    if err := ReflectorPrompt.Prompter("Enter your county of residence: ", &county); err != nil {
        panic(err)
    }
    switch county {
    case "Eau Claire":
        return 0.005
    case "Dunn":
        return 0.004
    default:
        return 0
    }

    return 0
}

func RoundUp(number float64) float64 {
    return math.Ceil(number * 100) / 100
}

func main() {
    var state string
    var amount float64

    // Get user input
    if err := ReflectorPrompt.Prompter("Enter the bill amount ($): ", &amount); err != nil {
        panic(err)
    }

    if err := ReflectorPrompt.Prompter("Enter your State of Residence (XX): ", &state); err != nil {
        panic(err)
    }

    taxrate := taxRate(state)
    tax := amount * taxrate
    total := amount + tax
    total = RoundUp(total)

    // Output the results
    fmt.Printf("The total amount payable is %.2f\n", total)
}