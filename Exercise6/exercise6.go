/*  Create a program that determines how many years you have left until retirement
    and the year you can retire and display the output as follows:

    What is your current age: 25
    At what age would you like to retire: 65
    You have 40 years left until you can retire
    It's currently 2016, so you can retire in 2056

 */

package main

import (
    "time"
    "fmt"
)

func main(){
    // Get user input
    var currentAge int
    var retirementAge int

    fmt.Print("What is your current age?: ")
    fmt.Scanf("%d ", &currentAge)

    fmt.Print("At what age would you like to retire?: ")
    fmt.Scanf("%d", &retirementAge)
    // Get difference in years between retirement age and current age
    deltaAge := retirementAge - currentAge

    // Add difference to current year to obtain final result
    currentYear := time.Now().Year()
    retirementYear := time.Now().AddDate(deltaAge,0,0).Year()

    // Print results
    fmt.Printf("You have %d years left until you can retire\n", deltaAge)
    fmt.Printf("It's currently %d, so you can retire in %d\n",currentYear, retirementYear)
}
