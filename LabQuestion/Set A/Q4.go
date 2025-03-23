//4. WAP in go Language to print address of a variable. 
package main

import "fmt"

func main() {
    var number int = 42
    fmt.Println("Value of number:", number)
    fmt.Println("Address of number:", &number)
}
