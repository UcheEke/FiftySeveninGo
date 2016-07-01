package main

import (

    "strings"
    "fmt"
    "../../../Projects/FiftySeven/ReflectorPrompt"
)

const (
    letters = "abcdefghijklmnopqrstuvwxzyABCDEFGHIJKLMNOPQRSTUVWXYZ"
    numbers = "0123456789"
    specials = "!£$%^&*()|/\\#@:;[]{}+=-_¬`?><.,"
)

func VeryWeakPassword(password string) bool {
    return (len(password) <= 8 &&
    strings.ContainsAny(password, numbers) &&
    !strings.ContainsAny(password, letters) &&
    !strings.ContainsAny(password, specials))
}

func WeakPassword(password string) bool {
    return (len(password) <= 8 &&
    strings.ContainsAny(password, letters) &&
    !strings.ContainsAny(password, numbers) &&
    !strings.ContainsAny(password, specials))
}

func ModeratePassword(password string) bool {
    return (len(password) >= 8  &&
    strings.ContainsAny(password, letters) &&
    !strings.ContainsAny(password, numbers) &&
    !strings.ContainsAny(password, specials)) ||
    (len(password) >= 8  &&
    strings.ContainsAny(password, numbers) &&
    !strings.ContainsAny(password, letters) &&
    !strings.ContainsAny(password, specials))
}

func StrongPassword(password string) bool {
    return len(password) >= 8 &&
    strings.ContainsAny(password,letters) &&
    strings.ContainsAny(password, numbers) &&
    !strings.ContainsAny(password, specials)
}

func VeryStrongPassword(password string) bool {
    return len(password) >= 8 &&
    strings.ContainsAny(password, letters) &&
    strings.ContainsAny(password, numbers) &&
    strings.ContainsAny(password, specials)
}

func PasswordChecker(password string) {
    switch {
    case VeryWeakPassword(password):
        fmt.Printf("The password \"%s\" is very weak! Please change it!\n", password)
    case WeakPassword(password):
        fmt.Printf("The password \"%s\" is weak. Adding a few more varied characters may help!\n", password)
    case ModeratePassword(password):
        fmt.Printf("The password \"%s\" is moderate. Adding a few more varied characters may help!\n", password")
    case StrongPassword(password):
        fmt.Printf("The password \"%s\" is strong, but there is room for improvement. Add some non-alphanumerics\n", password)
    case VeryStrongPassword(password):
        fmt.Printf("The password \"%s\" is very strong. But don't be complacent. Change it often!\n", password)
    }
}

func main(){
    var password string
    if err := ReflectorPrompt.Prompter("Please enter a password: ", &password); err != nil {
        panic(err)
    }

    PasswordChecker(password)
}