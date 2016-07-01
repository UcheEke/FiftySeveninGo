/*
    Write a program that will help you to determine how many months it will
    take to pay off a credit card balance. The program should ask the user
    to enter the balance of a credit card and the APR of the card. The program
    should then return the number of months needed.

    The formula is given by

    n = -1/30 x (log(1 +b/p(1-(1+i)30)) / log(1 + i)

    where n is the number of months
    i is the daily rate i.e. APR / 365
    b is the balance
    p is the monthly payment

    Constraints:

    Prompt for the card's APR. Do the division internally
    Prompt for the APR as a percentage not a decimal
    Round all fractions of a cent up to the nearest cent
 */

package main

import (
    "fmt"
    "../../../Projects/FiftySeven/ReflectorPrompt"
    "math"
)

func roundUpCents(x float64) float64 {
    return math.Ceil(x * 100) / 100
}

func monthsToPay(balance, APR, payment float64) float64 {
    i := APR / 365
    months := (-1. / 30.) * math.Log10(1 + (balance / payment) * (1 - math.Pow(1 + i,30))) / math.Log10(1 + i)
    return months
}

func main(){
    var balance, payment, APR float64
    var vars = []*float64{&balance, &payment, &APR}
    var prompts = []string{ "Enter the balance ($): ",
                            "How much is your monthly payment ($): ",
                            "What is the APR (%)?: ",
    }

    // Get user input
    for i,v := range vars {
        if err := ReflectorPrompt.Prompter(prompts[i],v); err != nil {
            panic(err)
        }
    }

    // Clean up data
    APR = APR / 100
    balance = roundUpCents(balance)
    payment = roundUpCents(payment)

    months := monthsToPay(balance, APR, payment)
    months = math.Ceil(months)
    fmt.Printf("It will take you %.0f months to pay off this credit card\n", months)
}