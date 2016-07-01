/*
    Create a program that takes a product name as input and
    retrieves the current price and quantity for that product.
    The product data is in a data file in the JSON format

    Print out the product name, price and quantity if the product is found
    If no product matches the search, state that no product was found and
    start over
 */

package main

import (
    "os"
    "bufio"
    "fmt"
    "io/ioutil"
    "encoding/json"
    "strings"
    "../../../Projects/FiftySeven/ReflectorPrompt"
    "log"
)

type ProductList struct {
    Products []Product `json:"products"`
}

type Product struct {
    Name string `json:"name"`
    Price float64 `json:"price"`
    Quantity int `json:"qty"`
}

func PopulateProductList() (ProductList, error) {
    fd, err := os.Open("products.json")
    if err != nil {
        fmt.Println("Error reading from file 'products.json'")
        return ProductList{}, err
    }

    var pl ProductList
    defer fd.Close()

    r := bufio.NewReader(fd)
    if b, err := ioutil.ReadAll(r); err == nil {
        json.Unmarshal(b, &pl)
    } else {
        return ProductList{}, err
    }
    return pl, nil
}

func ExportProductList(pl ProductList) error {
    b, err := json.Marshal(pl)
    if err !=  nil {
        return err
    }

    fp, err := os.Create("products.json")
    if err != nil {
        return err
    }

    fp.Write(b)
    return nil
}

func NewProduct() Product {
    var p Product

    var name string
    var price float64
    var qty int

    if err := ReflectorPrompt.Prompter("Name of Product: ",&name); err != nil {
        log.Fatal(err)
    }

    if err := ReflectorPrompt.Prompter("Price per unit ($): ", &price); err != nil {
        log.Fatal(err)
    }

    if err := ReflectorPrompt.Prompter("Quantity available (units): ", &qty); err != nil {
        log.Fatal(err)
    }

    p.Name = name
    p.Price = price
    p.Quantity = qty

    return p
}

func main(){
    var name string
    var results = make([]Product, 0)
    // Get product name from user
    fmt.Print("Enter the product name: ")
    fmt.Scanf("%s", &name)
    name = strings.TrimSpace(name)
    name = strings.ToLower(name)

    // Populate the product list from JSON data source
    plist, _ := PopulateProductList()

    // Iterate through product list searching for name
    // (Expand to use regular expressions eventually)
    for _, p := range plist.Products {
        n := strings.ToLower(p.Name)
        if name == n || strings.Contains(n, name) {
            results = append(results, p)
        }
    }

    if len(results) == 0 {
        fmt.Println("No product matches search term.")
        var resp string
        for {
            if err := ReflectorPrompt.Prompter("Would you like to add product to listing (y/n)? ", &resp); err != nil {
                log.Fatal(err)
                os.Exit(1)
            }

            resp = strings.ToLower(resp)
            resp = strings.TrimSpace(resp)
            if resp[0] == 'y' {
                newProduct := NewProduct()
                plist.Products = append(plist.Products, newProduct)
                ExportProductList(plist)
                break
            } else {
                os.Exit(0)
            }
        }
    } else {
        fmt.Printf("%10s | %10s | %10s\n", "Name", "Price", "Quantity")
        for _, r := range results {
            fmt.Printf("%10s | %10.2f | %10d\n", r.Name,r.Price, r.Quantity)
        }
    }
}

