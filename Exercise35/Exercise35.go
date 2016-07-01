/*
    Prompt for names. Collect the names into an array. When the user enters a blank line stop collecting names
    and pick a random winner from the list. Display the winner and exit
 */

package main

import (
    "../../../Projects/FiftySeven/ReflectorPrompt"
    "fmt"
    "math/rand"
    "time"
    "strings"
)

func pickAWinner(names []string) string {
    l := len(names)
    if l > 0 {
        rand.Seed(time.Now().UnixNano())
        return names[rand.Intn(len(names))]
    }

    return "NOBODY!! \n\n(PS: You need to enter at least one name!)"
}

func collectNames() []string {
    var name string
    var names []string = make([]string, 0)

    for {
        if err := ReflectorPrompt.Prompter("Enter a name: ", &name); err != nil {
            panic(err)
        } else {
            name = strings.TrimSpace(name)
            if name == "" {
                break
            } else {
                names = append(names, name)
            }
        }
    }
    return names
}

func start(){
    fmt.Println(
        "Enter the names of the contestants and\nthe computer will pick a random winner",
        "\n(Enter a blank line to stop)\n",
    )
}

func printEllipsis(){
    for i := 0; i < 8; i++ {
        if i == 4 {
            fmt.Print(" (drum roll) ")
        } else {
            fmt.Print(".")
        }
        time.Sleep(time.Millisecond * 250)
    }
}

func main(){
    start()
    names := collectNames()
    fmt.Print("And the winner is ")
    printEllipsis()
    fmt.Printf(". %s\n", pickAWinner(names))
}