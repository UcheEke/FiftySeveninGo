/*
    Create a program that compares two strings and determines if
    the two strings are anagrams. The program should prompt for
    both input strings and display the output as shown in the example
    that follows

    Enter two strings and I'll tell you if they are anagrams:
    Enter the first string: amble
    Enter the second string: blame
    "blame" and "amble" are anagrams

    Constraints:

    Implement a function named isAnagram(word1, word2) that returns a boolean
    Check that both words are the same length

 */

package main

import (
    "../../../Projects/FiftySeven/ReflectorPrompt"
    "fmt"
    "strings"
    "os"
)

func isAnagram(w1, w2 string) bool {
    for _,letter := range w1 {
        if !strings.Contains(w2, string(letter)) {
            return false
        }
    }
    return true
}

func isSameLength(w1,w2 string) bool {
    return len(w1) == len(w2)
}

func main() {
    var words = make([]string, 2,2)
    var prefix = []string{"first", "second"}

    fmt.Println("Enter two strings and I'll tell you if they are anagrams:")

    for i, _ := range words {
        if err := ReflectorPrompt.Prompter(fmt.Sprintf("Enter the %s word: ", prefix[i]), &words[i]); err != nil {
            panic(err)
        }
    }

    if !isSameLength(words[0], words[1]) {
        fmt.Println("Words must be of equal length")
        os.Exit(0)
    }

    if isAnagram(words[0],words[1]) {
        fmt.Printf("\"%s\" is an anagram of \"%s\"\n", words[0], words[1])
    } else {
        fmt.Printf("\"%s\" is not an anagram of \"%s\"\n", words[0], words[1])
    }
}