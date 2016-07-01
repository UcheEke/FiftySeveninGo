/*
    Write a program to evenly divide pizzas. Prompt for
    the number of people, the number of pizzas, and the number
    of slices per pizza. Ensure that the number of pieces comes out
    even. Display the number of pieces of pizza each person should get.
    If there are leftovers, show the number of leftover pieces

    Example:
    How many people? 8
    How many pizzas do you have? 2
    Number of slices per pizza: 8

    8 people with 2 pizzas
    Each person gets 2 pieces of pizza
    There are 0 leftover pieces
 */

package main

import (
    "fmt"
)

func isEven(x int) bool {
    return x % 2 == 0
}

func main() {

    var (
        numPeople int
        numPizzas int
        numSlices int
    )

    // Obtain user input
    fmt.Print("How many people?: ")
    fmt.Scanf("%d ", &numPeople)
    fmt.Print("How many pizzas?: ")
    fmt.Scanf("%d ", &numPizzas)
    fmt.Print("How many slices per pizza?: ")
    fmt.Scanf("%d ", &numSlices)

    // Calculate the total number of slices
    slices := numPizzas * numSlices

    // Find the rounded down quotient of total slices / 2 * number of people
    quotient := slices / numPeople

    // Find remainder
    remainder := slices % numPeople

    // Print results
    fmt.Printf("Number of slices/person = %d\n", quotient)
    fmt.Printf("Number of slices remaining = %d\n", remainder)
}
