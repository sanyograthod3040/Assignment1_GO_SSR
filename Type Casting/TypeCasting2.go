package main

import "fmt"

func main() {
	var floatVal float64 =42.25
	var intVal int=int(floatVal)
	fmt.Println("float value", floatVal)
	fmt.Println("Converted int value",intVal)
}
