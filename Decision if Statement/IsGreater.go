package main

import "fmt"

func main() {
    var num int
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)

    if num > 100 {
        fmt.Println("The number is greater than 100.")
    }
}
