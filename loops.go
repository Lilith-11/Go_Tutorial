package main
import "fmt"
func main(){
	for i:=1;i<5;i++{
		fmt.Println(i)
	}
	var x int =1
	for x<5 {
		fmt.Println(x)
		x++
	}
	for {
		fmt.Println("hello world!")
		break
	}
}