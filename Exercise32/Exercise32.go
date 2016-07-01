/*
    Write a "Guess the Number" game that has three levels of
    difficulty. The first level of difficulty would be a number b/w
    1 and 10, the next b/w 1 and 100 and level 3 b/w 1 and 1000

    Prompt for the difficulty level, and then start the game. The computer picks
    a random number in that range and prompts the player to guess that number
    Each time the player guesses the computer should give a hint as to whether the
    number is too high of too low. The computer should also keep track of the number
    of guesses. Once the player guesses the correct number the computer should present
    the number of guesses and ask if the player would like to try again

    This will need to be revised
 */

package main

import (
    "math/rand"
    "time"
    "fmt"
    "os"
    "strings"
)

func rangeFromLevel(level int) int {
    switch (level) {
    case 1:
        return 10
    case 2:
        return 100
    case 3:
        return 1000
    default:
        return 10
    }
}

func testGuess(guess, target int) int {
    switch {
    case guess == target:
        return 0
    case guess < target:
        return 1
    case guess > target:
        return 2
    }
    return 3
}

func initGame() (int, int) {
    var level, guesses int
    var loop bool = true
    fmt.Println("Guess the Number")
    fmt.Println()
    fmt.Println("Choose a level:")
    fmt.Println("1. Guess a number b/w 1 and 10")
    fmt.Println("2. Guess a number b/w 1 and 100")
    fmt.Println("3. Guess a number b/w 1 and 1000")
    fmt.Println()

    for loop {
        fmt.Print("What level would you like to play (0 to exit): ")
        fmt.Scanf("%d ", &level)
        switch level {
        case 0:
            fmt.Println("Goodbye!")
            os.Exit(0)
        case 1,2,3:
            fmt.Printf("\nStarting level %d ...\n",level)
            guesses = level + 4
            loop = false
        default:
            fmt.Printf("'%s' is not a valid option. Please try again!\n", string(level))
        }
    }
    return level, guesses
}

func setupNewLevel(level int) int {
    var low int = 1
    high := rangeFromLevel(level)
    fmt.Println("Guess a number b/w",low,"and",high)
    // Seed the random number generator
    rand.Seed(time.Now().UnixNano())
    var target int = low + rand.Intn(high)
    return target
}

func getNextGuess() int {
    var guess int
    for {
        fmt.Print("Enter your guess: ")
        fmt.Scanf("%d ", &guess)
        if guess > 0 {
            break
        } else {
            fmt.Println("Not a valid guess. Try again")
        }
    }
    return guess
}

func play(target, guesses int) {
    var lastGuess, guess int
    var gameOver bool = false
    var loop bool = true

    for loop {
        switch guesses {
        case 0:
            gameOver = true
            loop = false
        case 1:
            fmt.Println("Only one guess left...")
        default:
            fmt.Printf("Guesses: %d\tLast guess value: %d\n", guesses, lastGuess)
        }

        // Switch isn't directly interchangable with if-else blocks
        if loop == false {
            break
        }

        lastGuess = guess
        guess = getNextGuess()
        guesses -= 1

        // Using an if else block here will allow you to use the break statement to exit
        // the loop early. Or use boolean flags as I have here now
        switch testGuess(guess, target) {
        case 0:
            fmt.Println("Well done, you've guessed correctly!")
            loop = false
        case 1:
            fmt.Println("Too low!")
        case 2:
            fmt.Println("Too high!")
        }
    }

    if gameOver {
        fmt.Printf("You lost :( The number you were looking for was %d\n", target)
    }
}

func playAgain(){
    var resp string
    for {
        fmt.Print("Would you like to play again? (y/n): ")
        fmt.Scanf("%s ", &resp)
        resp = strings.TrimSpace(resp)
        if len(resp) == 1 && strings.ContainsAny(resp, "yn") {
            switch resp {
            case "y":
                return
            case "n":
                fmt.Println("Goodbye!")
                os.Exit(0)
            }
        }
    }

}

func main(){
    for {
        level, guesses := initGame()
        target := setupNewLevel(level)
        play(target, guesses)
        playAgain()
    }
}