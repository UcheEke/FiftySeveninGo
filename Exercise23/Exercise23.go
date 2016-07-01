package main

import (
    "fmt"
    "../../../Projects/FiftySeven/ReflectorPrompt"
    "strings"
    "os"
)

func checkInput(response string) (string, error) {
    response = strings.ToLower(response)
    response = strings.TrimSpace(response)
    if response == "yes" || response == "no" {
        return response, nil
    }
    return response, fmt.Errorf("Please answer 'yes' or 'no'")
}

func AskQuestion(question string) (string, error) {
    var str string
    for {
        if err := ReflectorPrompt.Prompter(question, &str); err != nil {
            return "", err
        }
        response, err := checkInput(str)
        if err != nil {
            fmt.Println(err.Error())
            continue
        } else {
            return response, nil
        }
    }
}

func PrintResponse(response string) {
    fmt.Println(response)
    os.Exit(0)
}

func handleErrors(err error) {
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }
}

func ExpertSystem() {
    response, err := AskQuestion("Is the car silent when you turn the key? ")
    handleErrors(err)
    switch response {
    case "yes":
        response, err := AskQuestion("Are the battery terminals corroded? ")
        handleErrors(err)
        switch response {
        case "yes":
            PrintResponse("Clean terminals and try starting again")
        case "no":
            PrintResponse("The cables may be faulty. Replace cables and start again")
        }
    case "no":
        response, err = AskQuestion("Does the car make a clicking noise? ")
        handleErrors(err)
        switch response {
        case "yes":
            PrintResponse("Sounds like the battery is dead. Replace the battery")
        case "no":
            response, err = AskQuestion("Does the car crank up but fail to start? ")
            handleErrors(err)
            switch response {
            case "yes":
                PrintResponse("Check spark plug connections")
            case "no":
                response, err = AskQuestion("Does the engine start and then die? ")
                handleErrors(err)
                switch response {
                case "yes":
                    response, err = AskQuestion("Does you car have fuel injection? ")
                    handleErrors(err)
                    switch response {
                    case "yes":
                        PrintResponse("Get it in for a service")
                    case "no":
                        PrintResponse("Check to ensure the choke is opening and closing")
                    }
                case "no":
                    PrintResponse("I'm all out of answers! Get it in for a service")
                }
            }
        }
    }
}

func main(){
    ExpertSystem()
}