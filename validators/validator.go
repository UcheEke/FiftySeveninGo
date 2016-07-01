package validators

import "fmt"

// GetUserFloat(prompt, failure string, isValid func(float64) bool) float64
// Takes a prompt string to prompt the user for an item, a failure string if the
// user input is invalid and a user defined validation function to test the response
func GetUserFloat(prompt, failure string, isValid func(float64) bool) float64 {
    var resp float64
    for {
        fmt.Print(prompt)
        fmt.Scanf("%f", &resp)
        if isValid(resp) {
            break
        } else {
            fmt.Println(failure)
        }
    }
    return resp
}

// GetUserInt(prompt, failure string, isValid func(int) bool) int
// Takes a prompt string to prompt the user for an item, a failure string if the
// user input is invalid and a user defined validation function to test the response
func GetUserInt(prompt, failure string, isValid func(int) bool) int {
    var resp int
    for {
        fmt.Print(prompt)
        fmt.Scanf("%d", &resp)
        if isValid(resp) {
            break
        } else {
            fmt.Println(failure)
        }
    }
    return resp
}

// GetUserString(prompt, failure string, isValid func(string) bool) string
// Takes a prompt string to prompt the user for an item, a failure string if the
// user input is invalid and a user defined validation function to test the response
func GetUserString(prompt, failure string, isValid func(string) bool) string {
    var resp string
    for {
        fmt.Print(prompt)
        fmt.Scanf("%s", &resp)
        if isValid(resp) {
            break
        } else {
            fmt.Println(failure)
        }
    }
    return resp
}

