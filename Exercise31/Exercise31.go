/*
    Create a program that prompts for
    your age and your resting heart rate. Use the Karvonen formula
    to determine the target heart rate based on a
    range oof intensities from 55% for 95%.

    Generate a table with the results as shown in the example output

    Formula:

    Target Heartrate = (((220-age) - restingHR)*intensity)+restingHR)

    Example:

    Resting Pulse: 65   Age: 22

    Intensity    | Rate
    -----------------------
    55%          | 138 bpm
    60%          | 145 bpm

    90%          | 185 bpm
    95%          | 191 bpm

 */

package main

import (
    "math"
    "fmt"
)

func KarvonenHR(age, rp, intensity float64) float64 {
    return (((220.0 - age) - rp)*intensity/100) + rp
}

func RoundUp(x float64) float64 {
    return math.Ceil(x)
}

func main() {
    var intensity float64 = 50
    var age, rp float64
    var vars = []*float64{&age, &rp}
    var prompts = []string{
        "Enter your age: ",
        "What is your resting heartbeat?: ",
    }

    // get user input
    for i, _ := range vars {
        fmt.Print(prompts[i])
        fmt.Scanf("%f ", vars[i])
    }

    // Print Table Headers
    fmt.Printf("\nKarvonen Heart Rate by Intensity\n")
    fmt.Printf("Resting Pulse: %.0f bpm   Age: %.0f\n",rp, age)
    fmt.Println("Intensity (%)\t| Rate (bpm)")
    fmt.Println("-----------------------------")

    for i:=0; i<10; i++ {
        fmt.Printf("%10.0f\t|\t%2.0f\n", intensity, KarvonenHR(age, rp, intensity))
        intensity += 5
    }
    fmt.Println()
}
