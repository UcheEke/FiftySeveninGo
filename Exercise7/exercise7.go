/*
    Create a program that calculates the area of a room. Prompt the user
    for the length and width of the room in feet. Then display the area
    in both square feet and square meters

    m2 = f2 x 0.09290304
 */

package main

import (
    "fmt"
)

const (
    sqf2m2 = 0.09290304
)

func main() {
    var (
        length float64
        width float64
    )

    // Prompt for user input
    fmt.Println("Input Dimensions:")

    fmt.Print("Enter the length (ft): ")
    fmt.Scanf("%f ", &length)

    fmt.Print("Enter the width (ft): ")
    fmt.Scanf("%f ", &width)

    // Calculate area
    areaft2 := length * width
    aream2 := areaft2 * sqf2m2

    // Output the results
    fmt.Printf("The area is %.2f sq.ft (%.2f sq m)\n", areaft2, aream2)
}


