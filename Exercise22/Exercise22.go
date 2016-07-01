package main

import (
	"../../../Projects/FiftySeven/ReflectorPrompt"
	"fmt"
)

func findMax(numbers []int) int {
	var max int
	max = numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	return max
}

func main() {
	var numbers = make([]int, 3, 3)
	var prefix = []string{"first", "second", "third"}
	// Get user input
	for i, _ := range numbers {
		if err := ReflectorPrompt.Prompter(fmt.Sprintf("Enter the %s integer: ", prefix[i]), &numbers[i]); err != nil {
			panic(err)
		}
	}
	max := findMax(numbers)
	fmt.Printf("The largest number is %d\n", max)
}
