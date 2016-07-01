package prompter

/*
    Ultimately, I would like to produce a function of the form:
    prompt(statement str, variable interface{})

    where the calling process provides a pointer to the var 'variable'
    The reflect package then interrogates the variable and determines
    what the format string for fmt.Scanf() will be, placing the user input
    prompted by the given statement into the variable address.

    The "laws of reflection" blogpost advises that we should avoid using the
    reflect package unless strictly necessary, which may imply that the prompter in this file may be
    the idiomatic way to go! But it would be a useful exercise nonetheless. ReflectorPrompt has
    been made as a trial run

*/

import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

type Prompt struct {}

func (p *Prompt) PromptInt(prompt string) int {
    result := new(int)// creates a pointer to the type
    fmt.Printf(prompt)
    fmt.Scanf("%d ", result)
    return *result
}

func (p *Prompt) PromptFloat(prompt string) float64 {
    result := new(float64)
    fmt.Printf(prompt)
    fmt.Scanf("%f ", result)
    return *result
}

func (p *Prompt) PromptStr(prompt string) string {
    result := new(string)
    fmt.Printf(prompt)
    fmt.Scanf("%s ", result)
    return *result
}

func (p *Prompt) PromptStrs(prompt string) string {
    fmt.Print(prompt)
    sc := bufio.NewReader(os.Stdin)
    result, err := sc.ReadString('\n')
    if err != nil {
        panic(err)
    }
    return strings.TrimSpace(result)
}