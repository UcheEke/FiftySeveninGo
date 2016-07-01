/*
    Magic 8 ball simulator
 */

package main

import (
    "math/rand"
    "time"
    "../../../Projects/FiftySeven/ReflectorPrompt"
    "fmt"
    "strings"
)

func createResponses() []string {
    var responses = []string{
        "Yes",
        "Sure!",
        "Maybe",
        "Certainly!",
        "Maybe not",
        "Sure, why not?!",
        "Apparently so!",
        "Possibly, it's all a little unclear",
        "Not so sure",
        "Nuh uh!",
        "Most definitely",
        "Not really!",
        "It wouldn't surprise me!",
        "Well, that would be a surprise!",
        "No",
        "Ask again later",
        "Most certainly not!",
        "Nope",
        "Um, nah!",
        "Seriously?! Erm....nah!",
        "Definitely",
        "Without hesitation, no!",
        "Without reservation, yes!",
    }
    return responses
}

func generateResponse(responses []string) string {
    i := len(responses)
    rand.Seed(time.Now().UnixNano())
    response := responses[rand.Intn(i)]
    return response
}

func getUserQuestion() string {
    var question string
    for {
        if err := ReflectorPrompt.Prompter("Ask me a question: ", &question); err != nil {
            fmt.Println(err)
            continue
        } else {
            break
        }
    }

    return strings.TrimSpace(question)
}

func askAnotherQuestion() bool {
    var resp string
    for {
        if err := ReflectorPrompt.Prompter("Do you want to ask another question? (y/n): ", &resp); err != nil {
            fmt.Println(err)
            continue
        } else {
            if len(strings.TrimSpace(resp)) == 1 && strings.ContainsAny(resp, "yn") {
                switch resp {
                case "y":
                    return true
                case "n":
                    return false
                }
            } else {
                fmt.Println("Please answer 'y' or 'n'")
                continue
            }
        }
    }
    return false
}

func main(){
    var play bool = true
    responses := createResponses()
    for play {
        getUserQuestion()
        response := generateResponse(responses)
        fmt.Println(response)
        play = askAnotherQuestion()
    }
}