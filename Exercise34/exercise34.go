/*
    Create a small program that contains a list of employee names
    Print out the list of names when the program runs for the first time
    Prompt for an employee name and remove that specific name from the list of names
    Display the remaining employees, each on its own line
 */

package main

import (
    "fmt"
    "../../../Projects/FiftySeven/ReflectorPrompt"
    "strings"
)

func displayEmployees(employees []string) {
    // Print out list of names
    fmt.Println("There are",len(employees),"employees\n")
    for _, name := range employees {
        fmt.Println(name)
    }
}

func main(){
    var eName string
    var employees = []string{
        "John Smith",
        "Jackie Jackson",
        "Chris Jones",
        "Amanda Cullen",
        "Jeremy Goodwin",
    }

    tempArr := make([]string,0,len(employees))
    displayEmployees(employees)

    // Prompt for an employee name to remove
    if err := ReflectorPrompt.Prompter("Enter an employee name to remove: ", &eName);err != nil {
        panic(err)
    }

    // Remove from list if present, print error message otherwise
    eName = strings.TrimSpace(eName)

    for _ ,name := range employees {
        if eName == name {
            continue
        } else {
            tempArr = append(tempArr, name)
        }
    }

    employees = tempArr

    // Display remaining names
    displayEmployees(employees)
}