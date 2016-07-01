/*
   Write a program that prompts for the response times from a website in milliseconds
   It should keep asking for values until the user enters "done"

   The program should print the average time (mean), the minimum time ,
   the maximum time, and the standard deviation

   To compute the average:
   1. Compute the sum
   2. Divide the sum by the number of values

   To compute the standard deviation:
   1. Calculate the difference from the mean for each number and square it
   2. Compute the mean of the squared values
   3. Take the square root of the mean

*/

package main

import (
	"../../../Projects/FiftySeven/ReflectorPrompt"
	"fmt"
	"math"
)

func mean(numbers ...float64) float64 {
	var sum float64
	for _, num := range numbers {
		sum += num
	}
	return sum / float64(len(numbers))
}

func max(numbers ...float64) float64 {
	var max float64 = numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	return max
}

func min(numbers ...float64) float64 {
	var min float64 = numbers[0]
	for _, num := range numbers {
		if num < min {
			min = num
		}
	}
	return min
}

func sDev(numbers ...float64) float64 {
	var sqVals = make([]float64, 0, len(numbers))
	mu := mean(numbers...)
	for _, num := range numbers {
		del := num - mu
		sqVals = append(sqVals, (del * del))
	}
	return mean(math.Sqrt(mean(sqVals...)))
}

func calculateOutliers(sd float64, numbers ...float64) []float64 {
    var outliers []float64 = make([]float64, 0)
    if sd <= 0 {
        return outliers
    }

    // Outliers are more than 2 * standard deviations away from mean
    mu := mean(numbers...)
    for _, num := range numbers {
        diff := num - mu
        if math.Abs(diff / sd) > 2 {
            outliers = append(outliers, num)
        }
    }

    return outliers
}

func collectData() []float64 {
	var data []float64 = make([]float64, 0)
	var dp float64
	// Explain
	fmt.Println("Enter the response times (-1 for 'done'): ")
	// Collect data:
	for {
		if err := ReflectorPrompt.Prompter("Response time (ms): ", &dp); err != nil {
			panic(err)
		} else {
            if dp < 0 {
                break
            } else {
                data = append(data, dp)
            }
		}
	}
	return data
}

func displayStatistics(data ...float64) {
	// Mean
	fmt.Printf("The mean response time is %.2f ms\n", mean(data...))
	// Min
	fmt.Printf("The minimum response time was %.2f ms\n", min(data...))
	// Max
	fmt.Printf("The maximum response time was %.2f ms\n", max(data...))
	// SDev
    sd := sDev(data...)
    fmt.Printf("The Standard Deviation for the distribution is %.2f ms\n", sd)
    outliers := calculateOutliers(sd, data...)
    fmt.Print("The outliers are: ")
    fmt.Println(outliers)
}

func main() {
	data := collectData()
	displayStatistics(data...)
}
