package main
import "fmt"
func main(){
	defer fmt.Println("defer line")
	fmt.Println("before panic")
	panic("Something went wrong!")
    fmt.Println("This will not be printed")


}