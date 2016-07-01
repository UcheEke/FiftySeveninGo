/*
    Password Creator:

    Enter the number of the characters for your password
    Enter the number of special characters
    Enter the number of digits to use

    Generate the password
    Print the password
 */

package main

import "fmt"

const (
    digits = "1234567890"
    letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    specials = "!$£%^&*()-_=+{}[];@#~?/><.|"
)

func generatePassword(length,numS,numD int) (string, error) {
    var password []byte

    if length < 0 {
        return "", fmt.Errorf("Length must be positive!")
    }
    if numS > length || numD > length {
        return "", fmt.Errorf("# of optional characters must not exceed password length")
    }
    if numS + numD > length {
        return "", fmt.Errorf("Sum of optional characters must not exceed password length")
    }

    numL := length - (numS + numD)

    // Add Specials
    for i:= 0; i<numS; i++ {

    }
}