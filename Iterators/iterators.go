package main

import "fmt"

type iter struct {
    value int
    count int
}

func (i *iter) next() {
    i.value += i.count
    i.count += 1
}

func main() {
    I := iter{}
    for j := 0; j < 5; j++ {
        I.next()
        fmt.Println("Triangular number",I.count,"is", I.value)
    }
}