package main

import "fmt"

func main() {
    var n, sum int
    fmt.Print("Enter a number: ")
    fmt.Scan(&n)

    i := 1
    for i <= n {
        sum += i
        i++
    }
    fmt.Println("Sum of first", n, "natural numbers is", sum)
}
