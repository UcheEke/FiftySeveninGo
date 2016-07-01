/*
    Create a program to calculate the body mass index (BMI) for
    a person using the person's height in inches and weight in
    pounds. The program should prompt for weight and height

    Calculate the BMI using the following formula:

    bmi = 703 * (weight / height**2)

    a bmi b/w 18.5 and 25 is considered normal and overweight or underweight if
    above or below that range respectively

 */

package main

import (
    "../../../Projects/FiftySeven/ReflectorPrompt"
    "fmt"
)

func bmi(h, w float64) float64 {
    return 703 * (w / (h * h))
}

func diagnose(bmi float64) int {
    switch {
    case bmi < 18.5:
        return 0
    case bmi >= 18.5 && bmi <= 25:
        return 1
    case bmi > 25:
        return 2
    }
    return -1
}

func main() {
    // Create data structures and input variables
    var diagnosis map[int]string = map[int]string{
        -1 : "An error occurred!",
        0  : "Underweight",
        1  : "Normal",
        2  : "Overweight",
    }

    var weight, height float64

    // Get user input
    if err := ReflectorPrompt.Prompter("Enter your weight (lbs): ", &weight); err != nil {
        panic(err)
    }
    if err := ReflectorPrompt.Prompter("Enter your height (inches): ", &height); err != nil {
        panic(err)
    }

    // Calculate bmi
    val := bmi(height, weight)

    // Test bmi
    diag := diagnose(val)

    // Output results
    fmt.Printf("Your BMI is given as %.2f, which is considered to be '%s'\n", val, diagnosis[diag])
}