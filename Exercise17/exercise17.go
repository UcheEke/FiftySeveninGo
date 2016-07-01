/*
    Create a program that prompts for your weight, gender, number
    of drinks, the amount of alcohol by volume of the drinks consumed
    and the amount of time since your last drink. Calculate you blood alcohol
    content using this formula

    Blood Alc. Content = (5.14A/Wr)-0.15H

    A = total alcohol consumed in oz
    W = body weight in lbs
    r = alcohol distribution ratio
        male = 0.73, female = 0.66
    H = number of hours since last drink

    Display whether it is legal to drive by comparing bac value to
    0.08.

    Constraint: Prevent user from entering non numeric values
 */

package main

import (
    "fmt"
    "strings"
)

const (
    rM = 0.73
    rF = 0.66
)

func BAC(A, W, r, H float64) float64 {
    return (5.14 * A / (W * r)) - 0.15 * H
}

func main() {
    var (
        A, W, n, abv, H float64
        gender string
        params []float64
        floater float64
    )

    params = []float64{W,n,abv,H}
    questions := []string{
        "What is your current weight (lbs)? ",
        "Number of fl.oz consumed: ",
        "Average abv of drinks (%vol): ",
        "Time since your last drink (hours)?: ",
    }

    for i := 0; i < len(params); i++ {
        fmt.Print(questions[i])
        fmt.Scanf("%f ", &floater)
        params[i] = floater
    }

    fmt.Printf("%v\n",params)

    fmt.Print("What is your biological gender? (m/f)? ")
    fmt.Scanf("%s ", &gender)
    gender = strings.TrimSpace(gender)
    gender = strings.ToLower(gender)
    r := rF
    if gender == "male" {
        r = rM
    }

    // Calculate A
    W = params[0]
    n = params[1]
    abv = params[2]
    H = params[3]

    A = n * abv

    bac := BAC(A,W,r,H)
    fmt.Printf("Your blood alcohol level is %.2f\n", bac)
    if bac >= 0.08 {
        fmt.Println("You are currently not legal to drive")
        startH := H
        for ;; H++ {
            if BAC(A,W,r,H) < 0.08 {
                fmt.Printf("You should be legal in %.0f hours time\n",H-startH)
                break
            }
        }
    } else {
        fmt.Println("You are legally able to drive")
    }
}