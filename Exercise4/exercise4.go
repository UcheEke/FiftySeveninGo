package main

import (
    "strings"
    "fmt"
)

func aoran(word string) string {
    if strings.ContainsAny(string(word[0]), "aeiou") {
        return "an"
    }
    return "a"
}
func main() {
    // Define variables
    keys := []string{"verb", "adj", "noun", "adverb"}
    var responses map[string]string = map[string]string{}
    var response string
    // Gain user input
    for _, key := range keys {
        fmt.Printf("Enter %s %s: ", aoran(key), key)
        fmt.Scanf("%s ", &response)
        responses[key] = response
    }

    // Concatenate result
    statement := fmt.Sprintf("\"Do you %s your %s %s %s?\"",
        responses["verb"],
        responses["adj"],
        responses["noun"],
        responses["adverb"],
    )
    // Output result
    fmt.Println(statement)
}
