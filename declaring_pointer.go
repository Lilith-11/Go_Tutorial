package main 
import "fmt"

func main() {
	var p *int

	fmt.Println("the value of p is", p)
	var x int=10
	var ptr*int=&x
	fmt.Println(x)
	fmt.Println(ptr)
	fmt.Println(&x)
	fmt.Println(*ptr)
}