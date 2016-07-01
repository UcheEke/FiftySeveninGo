package main

import (
    "../../../Projects/FiftySeven/ReflectorPrompt"
    "fmt"
)

func main(){
    var month int
    var months = map[int]string{
        1 : "January",
        2 : "February",
        3 : "March",
        4 : "April",
        5 : "May",
        6 : "June",
        7 : "July",
        8 : "August",
        9 : "September",
        10: "October",
        11: "November",
        12: "December",
    }

    if err := ReflectorPrompt.Prompter("Enter the number of the month: ", &month); err != nil {
        panic(err)
    }

    if v, ok := months[month]; !ok {
        fmt.Println("That month doesn't exist in the 12-month Calendar!")
    } else {
        fmt.Printf("The name of the month is %s\n", v)
    }
}