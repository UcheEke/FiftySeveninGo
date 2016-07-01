/*
   Create a program that computes simple interest. Prompt for the principal amount, the rate
   as a percentage, and the time and display the amount accrued

   The formula for simple interest is A=P(1+rt)

   Constraints:

   Prompt for the rate as a percentage, not as a number in range [0,1]
   Ensure that fractions of a cent are rounded up to the nearest penny
   Ensure the output is formatted as money

*/

package main

import (
	"fmt"
	"math"
)

func CalculateAmount(p, r, t float64) float64 {
	return p * (1 + r*t)
}

func fix(amount, dps float64) float64 {
	fixer := math.Pow(10, dps)
	return math.Floor(amount*fixer) / fixer
}

func RoundUp(amount float64) float64 {
	return math.Ceil(amount*100) / 100
}

func main() {
	var (
		principal float64
		rate      float64
		t         float64
	)

	// Prompt for user input
	fmt.Println("Simple Interest Calculator")
	fmt.Print("Enter the principal amount (£): ")
	fmt.Scanf("%f ", &principal)

	fmt.Print("Enter the annual rate: ")
	fmt.Scanf("%f ", &rate)

	fmt.Print("Enter the number of years: ")
	fmt.Scanf("%f ", &t)

	// Clean up input
	p := fix(principal, 2)
	r := fix(rate/100, 2)
	t = fix(t, 2)

	// Calculate amount accrued over time
	total := CalculateAmount(p, r, t)
	total = RoundUp(total)
	total = fix(total, 2)

	// Output the results
	fmt.Printf("Amount accrued on £%.2f after %.2f years @%.2f%% apr is £%.2f\n", p, t, rate, total)
}
