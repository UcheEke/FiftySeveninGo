/*
    Calculate the gallons of paint needed to paint the ceiling of a room
    Prompt for the length and width, and assume one gallon covers 350 sq ft.
    Display the number of gallons needed to paint the ceiling
 */

package main

import (
    "fmt"
    "math"
)

const (
    sqftPerGallon = 350
)

func main() {
    var (
        length float64
        width float64
    )

    // Prompt for user input (length, width)
    fmt.Printf("Dimensions:\n")
    fmt.Print("Enter the length (ft): ")
    fmt.Scanf("%f ", &length)

    fmt.Print("Enter the width (ft): ")
    fmt.Scanf("%f ", &width)

    // Calculate area
    area := length * width

    // divide by sqftPerGallon to get actual number of gallons
    gallons := area / sqftPerGallon

    // Round up the result to the nearest integer
    gallonsToPurchase := int(math.Ceil(gallons))

    // Display results
    fmt.Printf("Area of ceiling = %.2f sq.feet\n", area)
    fmt.Printf("Number of gallons required = %d\n", gallonsToPurchase)
}
