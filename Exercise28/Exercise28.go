/*
    Write a program that asks for 5 numbers and then prints out the sum of all
    numbers requested
 */

package main

import (
    "fmt"
)

func sum(nums ...float64) float64 {
    sum := 0.0
    for _, v := range nums {
        sum += v
    }

    return sum
}

func main(){
    var nums = make([]float64, 5, 5)

    for i, _ := range nums {
        fmt.Print("Enter a number: ")
        fmt.Scanf("%f ", &nums[i])
    }

    fmt.Printf("The sum of the numbers is %.2f\n", sum(nums...))
}