/*
    Write a program that prompts for a first name, last name, employee ID
    and ZIP code. Ensure that the input is valid according to the following
    rules:

    The first name must be filled in
    The last name must be filled in
    The first and last names must be at least 2 characters long
    The employee ID is of the format AA-1234
    The zip code must be a 5 digit number

    Constraints:

    Create a function for each validation type
    Use a single output statement to display the outputs
 */

package main

import (
    "fmt"
    "../../../Projects/FiftySeven/ReflectorPrompt"
    "strings"
)

const (
    caps = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    lowers = "abcdefghijklmnopqrstuvwxyz"
    numbers = "0123456789"
    specials = "\\|,.<>/?@#:;[]{}=+-_)(*&^%$£!`¬"
    letters = caps + lowers
    nonnumbers = letters + specials
    nonletters = numbers + specials
    noncaps = lowers + nonletters
)

func isValidName(name string) bool {
    return len(name) >= 2
}

func isValidID(id string) bool {
    return len(id) == 7 &&
    !strings.ContainsAny(id[:2], noncaps) &&
    id[2] == '-' &&
    !strings.ContainsAny(id[3:], nonnumbers)
}

func isValidZip(zip string) bool {
    return len(zip) == 5 &&
    !strings.ContainsAny(zip, nonnumbers)
}

func ValidateForm(fname, lname, id, zip string) {
    var output string

    if !isValidName(fname) {
        output += fmt.Sprintf("We can't accept \"%s\" as your first name. It needs to be at least two characters long\n", fname)
    }

    if !isValidName(lname){
    output += fmt.Sprintf("We can't accept \"%s\" as your last name. It needs to be at least two characters long\n", lname)
    }

    if !isValidID(id) {
        output += fmt.Sprintf("\"%s\" is not a valid Employee ID here!\n", id)
    }

    if !isValidZip(zip) {
        output += fmt.Sprintf("\"%s\" is not a valid zipcode (It should be numeric and 5 digits long)\n", zip)
    }

    if isValidName(fname) && isValidName(lname) && isValidID(id) && isValidZip(zip) {
        output = "The form is valid\n"
    }

    fmt.Println(output)
}

func main(){
    var fname, lname, id, zip string
    var vars = []*string{&fname, &lname, &id, &zip}
    var prompts = []string{
        "Enter your first name: ",
        "Enter your last name: ",
        "Your EmployeeID: ",
        "Your zipcode: ",
    }

    for i, v := range vars {
        if err := ReflectorPrompt.Prompter(prompts[i],v); err != nil {
            panic(err)
        }
    }

    ValidateForm(fname, lname, id, zip)
}
