package main
import "fmt"

const pi float64=3.1415
const language string ="Go"

func main(){
	var radius float64=5.0
	var area float64 = pi*radius*radius

	fmt.Println("constant pi",pi)
	fmt.Println("constant lang",language)
	fmt.Println("radius",radius)
	fmt.Println("Area of circle",area)

}