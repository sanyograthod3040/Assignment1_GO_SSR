package main

import "fmt"

func main() {
    var str string
    fmt.Print("Enter a string: ")
    fmt.Scanln(&str)

    for _, ch := range str {
        fmt.Printf("Character: %c, ASCII: %d\n", ch, ch)
    }
}
