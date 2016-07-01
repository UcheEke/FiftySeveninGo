/*
    Create a simple self checkout system. Prompt for the prices
    and quantities of three items. Calculate the subtotal of the items
    Then calculate the tax using a tax rate of 5.5%. Print out the line
     items with the quantity and total and then print out the subtotal, tax
     amount and total
 */

package main

import (
    "fmt"
    "math"
)

const (
    tax = 0.055
)

type Item struct {
    price float64
    qty int
}

func fix(val, sigf float64) float64 {
    t := math.Pow(10, sigf)
    return math.Floor(val * t) / t
}

func main(){
    // Items will be called item 1, item 2 ...
    receipt := make(map[int]Item)

    var (
        price float64
        qty int
    )

    // Prompt for 3 items
    for i := 1; i < 4; i++ {
        item := new(Item)
        fmt.Printf("Enter price of Item %d: ",i)
        fmt.Scanf("%f ", &price)
        item.price = fix(price, 2)
        fmt.Printf("Enter the quantity: ")
        fmt.Scanf("%d ", &qty)
        item.qty = qty

        receipt[i] = *item
    }

    // Calculate the subtotal
    var subtotal float64
    for _, item := range receipt {
        subtotal += (item.price * float64(item.qty))
    }

    subtotal = fix(subtotal, 2)

    // Calculate the tax @ 5.5%
    salesTax := subtotal * tax
    salesTax = fix(salesTax, 2)

    // Calculate the total
    total := subtotal + salesTax
    total = fix(total, 2)

    fmt.Printf("Subtotal : $%.2f\n", subtotal)
    fmt.Printf("Sales Tax @ 5.5%% : $%.2f\n", salesTax)
    fmt.Printf("Total : $%.2f\n", total)
}
