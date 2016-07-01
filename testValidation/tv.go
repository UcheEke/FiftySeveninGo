package main

import (
    "../../../Projects/FiftySeven/validators"
    "fmt"
)

func main() {
    resp := validators.GetUserInt(
        "Enter an value b/w 1 and 10: ",
        "Not a valid input.Please try again",
        func(r int) bool {
            return r > 0 && r < 10
        })

    fmt.Println("You entered",resp)
}
