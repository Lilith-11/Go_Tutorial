package main
import "fmt"
func main(){
	var x int =10
	fmt.Println("before x",x)
	changevalue(&x)
	fmt.Println("after x",x)
}
func changevalue(p *int){
	*p=30
}