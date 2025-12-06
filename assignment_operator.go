package main
import "fmt"
func main(){
	var x int
	x=10
	fmt.Println(x)
	x+=10
	fmt.Println("x+10",x)
	x-=10
	fmt.Println("x-10",x)
	x*=2
	fmt.Println("x*2",x)
	x/=2
	fmt.Println("x/2",x)
	x%=3
	fmt.Println("x%3",x)

}