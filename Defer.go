package main
import "fmt"
func main(){
	defer fmt.Println("this is performed last")
	fmt.Println("this is performed first")
}