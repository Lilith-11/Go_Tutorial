package main
import "fmt"
func main(){
	fmt.Println("misselaneous_operator")
	var x int = 10
	var w *int =&x
	fmt.Println(x)
	fmt.Println(&x)
	fmt.Println(w)
	fmt.Println(*w)

}