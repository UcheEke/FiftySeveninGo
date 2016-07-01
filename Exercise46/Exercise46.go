package main

import (
    "bufio"
    "os"
    "log"
    "strings"
    "fmt"
)

/*
    Create a program that reads in a file and counts the frequency of words within
    it. the construct a histogram displaying the words and the frequency, and display the histogram to
    the screen
 */

func cleanWord(w string) string {
    w = strings.Trim(w,";.,:!?'-&")
    w = strings.ToLower(w)
    w = strings.TrimSpace(w)
    return w
}

func processWords(words []string, hist map[string]int, count *int) {
    for _, word := range words {
        word = cleanWord(word)
        *count += 1
        if _, ok := hist[word]; !ok {
            hist[word] = 1
        } else {
            hist[word] += 1
        }
    }
}

func main() {
    var hist = make(map[string]int)

    fd, err := os.Open("macbeth.txt")
    if err != nil {
        log.Fatal(err)
    }

    sc := bufio.NewScanner(fd)

    var count int
    for sc.Scan() {
        line := sc.Text()
        words := strings.Fields(line)
        processWords(words, hist, &count)
    }

    fmt.Printf("%-20s\t%-7s\n", "Word", "Count")
    for k, v := range hist {
        fmt.Printf("%-20q\t%5d\n", k, v)
    }

    fmt.Println("Total word count:",count)
}
