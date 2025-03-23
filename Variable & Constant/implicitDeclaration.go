package main
import "fmt"

const pi =3.1415
const language  ="Go"

func main(){
	 radius :=7.0
	 area := pi*radius*radius

	fmt.Println("constant pi",pi)
	fmt.Println("constant lang",language)
	fmt.Println("radius",radius)
	fmt.Println("Area of circle",area)

}