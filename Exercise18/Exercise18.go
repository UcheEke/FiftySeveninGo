/*
    Create a program that converts temperatures from Fahrenheit to Celsius
    or vice verse. Prompt for the starting temperature. The program should
    prompt for the type of conversion and then perform the conversion
 */

package main

import (
    "../../../Projects/FiftySeven/ReflectorPrompt"
    "fmt"
)

func CtoF(degC float64) (degF float64) {
    return (degC * 9/5) + 32
}

func FtoC(degF float64) (degC float64) {
    return (degF - 32) * 5 / 9
}

func Convert(conversionType string, value float64) (float64, string) {
    if conversionType == "1" {
        return CtoF(value), "Fahrenheit"
    } else if conversionType == "2" {
        return FtoC(value), "Centigrade"
    }
    return 0, "ERR"
}

func main(){
    var degrees float64
    var convType string

    fmt.Println("Temperature Converter")

    // Get user input
    for {
        if err := ReflectorPrompt.Prompter("Enter the temperature value: ", &degrees); err != nil {
            panic(err)
        }
        fmt.Println("1. Celsius to Fahrenheit")
        fmt.Println("2. Fahrenheit to Celsius")
        if err := ReflectorPrompt.Prompter("What type of conversion do you wish to perform? ", &convType); err != nil {
            panic(err)
        }
        if convType == "1" || convType == "2" {
            break
        } else {
            fmt.Println("Unknown option entered. Please try again...")
        }
    }

    // Calculate the temperature
    temp, suffix := Convert(convType, degrees)

    // Output the results
    fmt.Printf("Temperature in %s is %.2f degrees\n", suffix, temp)
}