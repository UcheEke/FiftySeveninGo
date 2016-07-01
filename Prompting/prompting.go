package main

import (
    "../../../Projects/FiftySeven/prompter"
    "../../../Projects/FiftySeven/ReflectorPrompt"
    "fmt"
    "os"
)

func testPrompter(){
    p := new(prompter.Prompt)

    a := p.PromptInt("Give me an integer: ")
    fmt.Printf("The value %d is an %T\n", a,a)

    b := p.PromptFloat("Give me a floating point number: ")
    fmt.Printf("The value %.2f is a %T\n", b,b)

    s := p.PromptStr("Give me an string: ")
    fmt.Printf("The value '%s' is a %T\n", s,s)

    ss := p.PromptStrs("Give me a sentence: ")
    fmt.Printf("The value '%s' is a %T\n", ss,ss)
}

type MyVar int

func testReflectPrompter(){
    var number string
    var num2 MyVar
    if p := ReflectorPrompt.Prompter("Enter a string value: ", &number); p != nil {
        fmt.Println("\nError:", p)
        os.Exit(1)
    }

    fmt.Printf("The value of the string is '%s'\n", number)

    if p := ReflectorPrompt.Prompter("Enter another value: ", &num2); p != nil {
        fmt.Println("\nError:", p)
        os.Exit(1)
    }
}

func main() {
    testReflectPrompter()
}