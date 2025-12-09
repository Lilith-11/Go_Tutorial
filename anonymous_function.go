package main
import "fmt"
func main(){
	func(){
		fmt.Println("hello from the other side")
	}()
}