package main

import (
    "bufio"
    "os"
    "fmt"
    "strings"
)

func main(){
    reader := bufio.NewReader(os.Stdin)

    questions := map[string]string{
        "What is the quote? " : "",
        "Who said it? " : "",
    }

    for question , _ := range questions {
        fmt.Print(question)
        questions[question], _ = reader.ReadString('\n')
        questions[question] = strings.TrimSpace(questions[question])
    }

    fmt.Println(questions["Who said it? "] + " says, \"" + questions["What is the quote? "] + "\"")
}