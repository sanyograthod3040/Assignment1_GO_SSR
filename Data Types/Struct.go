package main
import "fmt"
 type Person struct{
	Name string
	Age int
 }

func main(){
	p:=Person{Name:"Sakshi",Age : 21}
	fmt.Println("personal details",p)
}