package main

import "fmt"

func main() {
	var char rune ='A'
	var str string=string(char)
	fmt.Println("Rune values", char)
	fmt.Println("Converted string value",str)

	var newRune rune = rune(str[0])
	fmt.Println("converted rune values",newRune)
}
