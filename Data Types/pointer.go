package main
import "fmt"

func main(){
	var num int =50
	var ptr *int=&num

fmt.Println("address of num ",ptr)	
fmt.Println("value of num through pointer",*ptr)
}