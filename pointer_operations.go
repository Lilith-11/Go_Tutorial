package main
import "fmt"
func main(){
	var x int =10
	var ptr *int=&x
	fmt.Println("x",x)
	*ptr=30
	fmt.Println("modified x ",x)
	fmt.Println("*ptr",*ptr)

}