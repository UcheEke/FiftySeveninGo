package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

/*  Create a program that prompts for an input string and
    displays output that shows the input string and the
    number of characters the string contains
 */

func main(){
    reader := bufio.NewReader(os.Stdin)
    var inputStr string
    // Display prompt for input string
    fmt.Print("Please enter a string of any length: ")
    // Get input string
    inputStr, _ = reader.ReadString('\n')
    // Prepare output: input string + number of characters
    outputStr := fmt.Sprintf("Input: %sNo. of Characters: %d", inputStr, len(strings.TrimSpace(inputStr)))
    // Display output
    fmt.Println(outputStr)
}