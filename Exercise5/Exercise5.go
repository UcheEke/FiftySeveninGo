package main

import "fmt"

func GetNumbers() []float64 {
    var num float64
    numbers := make([]float64, 2, 2)
    prefix := []string{"first", "second"}

    for i,_ := range numbers {
        fmt.Printf("What is the %s number? ",prefix[i])
        fmt.Scanf("%f ",&num)
        numbers[i] = num
    }
     return numbers
}

func main(){
    numbers := GetNumbers()
    num1 := numbers[0]
    num2 := numbers[1]

    fmt.Printf("%.1f + %.1f = %.1f\n",num1, num2, num1 + num2)
    fmt.Printf("%.1f - %.1f = %.1f\n",num1, num2, num1 - num2)
    fmt.Printf("%.1f x %.1f = %.1f\n",num1, num2, num1 * num2)
    fmt.Printf("%.1f / %.1f = %.1f\n",num1, num2, num1 / num2)
}
