package main

import (
	"fmt"
	"math"
)

func CompoundInterest(p, r, t, n float64) float64 {
	term := (1 + r/n)
	return p * math.Pow(term, n*t)
}

func RoundUp(amount float64) float64 {
	return math.Ceil(amount*100) / 100
}

func main() {
	var (
		p float64
		r float64
		t float64
		n float64
	)

	// Prompt for user input
	fmt.Println("Compound Interest Calculator: ")
	fmt.Print("Enter the amount: ")
	fmt.Scanf("%f ", &p)

	fmt.Print("Enter the rate: ")
	fmt.Scanf("%f ", &r)

	fmt.Print("Number of times compounded annually: ")
	fmt.Scanf("%f ", &n)

	fmt.Print("Number of years: ")
	fmt.Scanf("%f ", &t)

    // Clean up input
    rate := r / 100
	// Calculate the compound interest
	amount := CompoundInterest(p, rate, t, n)
	amount = RoundUp(amount)
	fmt.Printf("Principal amount: $%.2f\nRate %.2f%% compounded %.0f times per annum\nTotal accrued after %.0f years: $%.2f\n", p, r, n, t, amount)

}
