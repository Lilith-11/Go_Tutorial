package main
import "fmt"
func product(a,b int)int{
	return a*b
}
func main(){
	p:=product(2,3)
	fmt.Println("product is",p)
}