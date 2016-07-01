/*
    Using the 'rule of 72' to determine the approximate number of years one doubles
    their investment given a particular rate of interest, write a quick calculator that
    validity checks the rate to ensure that division is possible in the formula:

    n = 72 / r

    where   n = number of years
            r = rate of return stated
 */

package main

import (
    "fmt"
    "math"
)

func main(){
    var rate float64
    const epsilon = 0.0000001

    for {
        fmt.Print("Enter the stated rate of return: ")
        fmt.Scanf("%f", &rate)

        if rate > epsilon {
            break
        } else {
            fmt.Println("Invalid entry. Please try again")
        }
    }

    years := math.Ceil(72. / rate)

    fmt.Printf("You should expect to double your investment in %.0f years\n", years)
}
