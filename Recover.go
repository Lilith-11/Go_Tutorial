package main
import "fmt"
func main(){
	defer func(){
		if r:=recover();r!=nil{
			fmt.Println("recovered from panic",r)
		}
	}()
	fmt.Println("before panic")
	panic("something went wrong")
	fmt.Printf("this will not to be printed")

}